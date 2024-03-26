package main

import (
	"log"
	"os"

	"net/http"

	"github.com/asishshaji/notion-backup-go/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file %s", err)
	}
	client := &http.Client{Transport: &http.Transport{
		MaxIdleConns:       10,
		DisableCompression: true,
	}}

	notionApp := app.NewApp(os.Getenv("NOTION_TOKEN"), os.Getenv("NOTION_FILE_TOKEN"), os.Getenv("NOTION_SPACE_ID"), client)

	err = notionApp.ExportFromNotion(app.MardownExportType)
	if err != nil {
		log.Println(err)
	}
}
