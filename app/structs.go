package app

import (
	"encoding/json"
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

type GetTaskResult struct {
	ID        string `json:"id"`
	EventName string `json:"eventName"`
	Request   struct {
		SpaceID       string `json:"spaceId"`
		ExportOptions struct {
			ExportType string `json:"exportType"`
			TimeZone   string `json:"timeZone"`
			Locale     string `json:"locale"`
		} `json:"exportOptions"`
		ShouldExportComments bool `json:"shouldExportComments"`
	} `json:"request"`
	Actor struct {
		Table string `json:"table"`
		ID    string `json:"id"`
	} `json:"actor"`
	RootRequest struct {
		EventName string `json:"eventName"`
		RequestID string `json:"requestId"`
	} `json:"rootRequest"`
	Headers struct {
		IP                 string `json:"ip"`
		CityFromIP         string `json:"cityFromIp"`
		CountryCodeFromIP  string `json:"countryCodeFromIp"`
		Subdivision1FromIP string `json:"subdivision1FromIp"`
	} `json:"headers"`
	EqueuedAt int64 `json:"equeuedAt"`
	Status    struct {
		Type          string `json:"type"`
		PagesExported int    `json:"pagesExported"`
		ExportURL     string `json:"exportURL"`
	} `json:"status"`
	State string `json:"state"`
}

type GetTasksJSONResp struct {
	Results []GetTaskResult `json:"results"`
}

type ExportType string

const (
	MardownExportType ExportType = "markdown"
	HtmlExportType    ExportType = "html"
)

func MarshalRequest(r any) ([]byte, error) {
	marshalled, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}
