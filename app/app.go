package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type App struct {
	NOTION_TOKEN      string
	NOTION_FILE_TOKEN string
	NOTION_SPACE_ID   string
	Client            *http.Client
	TaskIds           []string
}

func NewApp(token, fileToken, spaceId string, httpClient *http.Client) *App {
	return &App{
		NOTION_TOKEN:      token,
		NOTION_FILE_TOKEN: fileToken,
		NOTION_SPACE_ID:   spaceId,
		Client:            httpClient,
	}
}

func (a *App) enqueueTask(exportType ExportType) (TaskJSONResp, error) {
	taskRequest := TaskJSONReq{
		T: Task{
			EventName: "exportSpace",
			Request: TaskRequest{
				SpaceId: a.NOTION_SPACE_ID,
				ExportOptions: ExportOptions{
					ExportType: string(exportType),
					TimeZone:   "America/New_York",
					Locale:     "en",
				},
				ShouldExportComments: false,
			},
		},
	}

	marshalledTaskRequest, err := MarshalRequest(taskRequest)
	if err != nil {
		return TaskJSONResp{}, err
	}

	req, err := http.NewRequest("POST", NOTION_API_ENQUEUE_URL, bytes.NewReader(marshalledTaskRequest))
	if err != nil {
		return TaskJSONResp{}, err
	}

	req.Header.Add("content-type", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "token_v2",
		Value: a.NOTION_TOKEN,
	})
	req.AddCookie(&http.Cookie{
		Name:  "file_token",
		Value: a.NOTION_FILE_TOKEN,
	})

	res, err := a.Client.Do(req)
	if err != nil {
		return TaskJSONResp{}, err
	}

	body, _ := io.ReadAll(res.Body)
	var taskResp TaskJSONResp
	err = json.Unmarshal(body, &taskResp)
	if err != nil {
		return TaskJSONResp{}, fmt.Errorf("error parsing enqueued task id : %s", err)
	}

	return taskResp, nil
}

// formats markdown and html
func (a *App) ExportFromNotion(exportType ExportType) error {
	taskResp, err := a.enqueueTask(exportType)
	if err != nil {
		return err
	}

	taskId := taskResp.TaskId
	log.Printf("enqueued task with id %s\n", taskId)
	a.TaskIds = append(a.TaskIds, taskId)

	return nil
}
