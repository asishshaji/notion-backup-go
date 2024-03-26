package app

import (
	"encoding/json"
	"log"
)

type ExportOptions struct {
	ExportType string `json:"exportType"`
	TimeZone   string `json:"timeZone"`
	Locale     string `json:"locale"`
}
type TaskRequest struct {
	SpaceId              string        `json:"spaceId"`
	ExportOptions        ExportOptions `json:"exportOptions"`
	ShouldExportComments bool          `json:"shouldExportComments"`
}
type Task struct {
	EventName string      `json:"eventName"`
	Request   TaskRequest `json:"request"`
}

type TaskJSONReq struct {
	T Task `json:"task"`
}

type TaskJSONResp struct {
	TaskId string `json:"taskId"`
}

type ExportType string

const (
	MardownExportType ExportType = "markdown"
	HtmlExportType    ExportType = "html"
)

func MarshalRequest(r any) ([]byte, error) {
	marshalled, err := json.Marshal(r)
	log.Printf("%s\n\n", string(marshalled))
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}
