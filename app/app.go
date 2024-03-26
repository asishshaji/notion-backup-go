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
	GetTaskResultCh   chan GetTaskResult
}

func NewApp(token, fileToken, spaceId string, httpClient *http.Client) *App {
	return &App{
		NOTION_TOKEN:      token,
		NOTION_FILE_TOKEN: fileToken,
		NOTION_SPACE_ID:   spaceId,
		Client:            httpClient,
		GetTaskResultCh:   make(chan GetTaskResult, 10),
	}
}
func (a *App) post(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing enqueued task id : %s", err)
	}

	return respBody, nil
}

func (a *App) getTasks(tasksIds []string) error {
	var tasksResponse GetTasksJSONResp

	var NOTION_API_GET_TASKS_URL string = fmt.Sprintf("%s/getTasks", NOTION_API_URL)
	taskIdsReq := struct {
		TaskIds []string `json:"taskIds"`
	}{
		TaskIds: tasksIds,
	}
	r, err := MarshalRequest(taskIdsReq)
	if err != nil {
		return fmt.Errorf("error marshalling getTasks request: %s", err)
	}
	resp, err := a.post(NOTION_API_GET_TASKS_URL, r)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(resp, &tasksResponse); err != nil {
		return err
	}

	for _, r := range tasksResponse.Results {
		a.GetTaskResultCh <- r
	}
	return nil
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

	var NOTION_API_ENQUEUE_URL string = fmt.Sprintf("%s/enqueueTask", NOTION_API_URL)

	resp, err := a.post(NOTION_API_ENQUEUE_URL, marshalledTaskRequest)
	if err != nil {
		return TaskJSONResp{}, err
	}
	var taskResp TaskJSONResp
	if err = json.Unmarshal(resp, &taskResp); err != nil {
		return TaskJSONResp{}, err
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

	// the tasks won't be ready yet, need to wait.
	// or keep polling if the task is ready
	a.getTasks(a.TaskIds)
	go a.processGetTaskResults()

	return nil
}

func (a *App) processGetTaskResults() {

	for range a.GetTaskResultCh {

	}
}
