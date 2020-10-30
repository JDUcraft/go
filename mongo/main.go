package main

import (
	models "./lib/support_thumbnail/infrastructure"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", models.HandleController(models.SupportThumbnailRepository{}))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

