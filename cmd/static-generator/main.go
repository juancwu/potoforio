package main

import (
	"context"
	"log"
	"os"

	"github.com/juancwu/potoforio/data"
	"github.com/juancwu/potoforio/views"
)

func main() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer f.Close()

	d := data.GetPageData()
	component := views.Index(d)

	err = component.Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}

	log.Println("Static site generated successfully to index.html")
}
