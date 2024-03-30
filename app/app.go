package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	NOTION_TOKEN      string
	NOTION_FILE_TOKEN string
	NOTION_SPACE_ID   string
	Client            *http.Client
	Quit              chan struct{}
	ExportUrls        chan string
}

func NewApp(token, fileToken, spaceId string, httpClient *http.Client) *App {
	a := &App{
		NOTION_TOKEN:      token,
		NOTION_FILE_TOKEN: fileToken,
		NOTION_SPACE_ID:   spaceId,
		Client:            httpClient,
		Quit:              make(chan struct{}, 2),
		ExportUrls:        make(chan string, 2),
	}

	return a
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
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing enqueued task id : %s", err)
	}

	return respBody, nil
}

func (a *App) getTasksStatus(taskId string) (GetTasksJSONResp, error) {
	var tasksResponse GetTasksJSONResp

	var NOTION_API_GET_TASKS_URL = fmt.Sprintf("%s/getTasks", NOTION_API_URL)
	taskIdsReq := struct {
		TaskIds []string `json:"taskIds"`
	}{
		TaskIds: []string{taskId},
	}
	r, err := MarshalRequest(taskIdsReq)
	if err != nil {
		return tasksResponse, fmt.Errorf("error marshalling getTasks request: %s", err)
	}
	resp, err := a.post(NOTION_API_GET_TASKS_URL, r)
	if err != nil {
		return tasksResponse, err
	}

	if err = json.Unmarshal(resp, &tasksResponse); err != nil {
		return tasksResponse, err
	}

	return tasksResponse, nil
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

	var NOTION_API_ENQUEUE_URL = fmt.Sprintf("%s/enqueueTask", NOTION_API_URL)

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

	go a.queryTaskStatus(taskResp.TaskId, exportType)

	return nil
}

func (a *App) queryTaskStatus(taskId string, exportType ExportType) {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			res, _ := a.getTasksStatus(taskId)
			state := res.Results[0].State
			if state == "success" {
				log.Println(res.Results[0].Status.ExportURL)
				go a.downloadExport(res.Results[0].Status.ExportURL, exportType)
				return
			} else if state == "in_progress" {
				log.Printf("%s polled, status %s\n", taskId, state)
			}
		case <-time.After(time.Second * 60):
			log.Println("shutting down")
			a.Quit <- struct{}{}
			return
		}
	}
}

func (a *App) downloadExport(url string, exportType ExportType) {
	log.Printf("url %s -> type %s", url, string(exportType))
	response, err := a.Client.Get(url)
	if err != nil {
		log.Println("Error while downloading ZIP file:", err)
		return
	}
	defer response.Body.Close()

	outputFile, err := os.Create(fmt.Sprintf("%s.zip", string(exportType)))
	if err != nil {
		log.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		log.Println("Error copying content to file:", err)
		return
	}

	log.Println("ZIP file downloaded successfully")
	a.Quit <- struct{}{}
}
