package main

import (
	"os"

	"net/http"

	"github.com/asishshaji/notion-backup-go/app"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// if err != nil {
	// 	log.Fatalf("error loading .env file %s", err)
	// }
	client := &http.Client{Transport: &http.Transport{
		MaxIdleConns:       10,
		DisableCompression: true,
	}}

	notionApp := app.NewApp(os.Getenv("NOTION_TOKEN"), os.Getenv("NOTION_FILE_TOKEN"), os.Getenv("NOTION_SPACE_ID"), client)
	exportTypes := []app.ExportType{app.HtmlExportType, app.MardownExportType}

	for _, t := range exportTypes {
		go notionApp.ExportFromNotion(t)
	}

	for range exportTypes {
		<-notionApp.Quit
	}
}
