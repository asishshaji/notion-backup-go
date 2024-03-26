package app

import "fmt"

const NOTION_API_URL = "https://www.notion.so/api/v3"

var NOTION_API_ENQUEUE_URL string = fmt.Sprintf("%s/enqueueTask", NOTION_API_URL)
