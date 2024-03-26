package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const NOTION_API_URL = "https://www.notion.so/api/v3"

type App struct {
	NOTION_TOKEN      string
	NOTION_FILE_TOKEN string
	NOTION_SPACE_ID   string
	Client            *http.Client
}

type ExportOptions struct {
	ExportType string `json:"exportType"`
	TimeZone   string `json:"timeZone"`
	Locale     string `json:"locale"`
}
type TaskRequest struct {
	SpaceId       string        `json:"spaceId"`
	ExportOptions ExportOptions `json:"exportOptions"`
}
type Task struct {
	EventName string      `json:"eventName"`
	Request   TaskRequest `json:"request"`
}

func (a *App) exportFromNotion() error {
	task := Task{
		EventName: "exportSpace",
		Request: TaskRequest{
			SpaceId: a.NOTION_SPACE_ID,
			ExportOptions: ExportOptions{
				ExportType: "markdown",
				TimeZone:   "America/New_York",
				Locale:     "en",
			},
		},
	}

	postUrl := fmt.Sprintf("%s/enqueueTask", NOTION_API_URL)
	marshalled, err := json.Marshal(task)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", postUrl, bytes.NewReader(marshalled))
	if err != nil {
		return err
	}

	req.Header.Add("Cookie",
		fmt.Sprintf("token_v2=%s; file_token=%s", a.NOTION_TOKEN, a.NOTION_FILE_TOKEN))
	res, err := a.Client.Do(req)
	if err != nil {
		return err
	}

	log.Println(res.Status)

	return nil
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file %s", err)
	}

	app := App{
		NOTION_TOKEN:      os.Getenv("NOTON_TOKEN"),
		NOTION_FILE_TOKEN: os.Getenv("NOTION_FILE_TOKEN"),
		NOTION_SPACE_ID:   os.Getenv("NOTION_SPACE_ID"),
		Client: &http.Client{Transport: &http.Transport{
			MaxIdleConns:       10,
			DisableCompression: true,
		}},
	}

	err = app.exportFromNotion()
	if err != nil {
		log.Println(err)
	}
}
