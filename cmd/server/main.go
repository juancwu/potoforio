package main

import (
	"log"
	"net/http"

	"github.com/juancwu/potoforio/data"
	"github.com/juancwu/potoforio/views"
)

func main() {
	// Serve static files (images, manifest, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Serve assets (compiled css)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Serve the Templ component
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d := data.GetPageData()
		component := views.Index(d)
		component.Render(r.Context(), w)
	})

	log.Println("Dev server starting on :8080 (Proxy to :9000 via Templ)")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
