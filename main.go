package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var embeddedFiles embed.FS

//go:embed views/*.html
var embeddedTemplates embed.FS

type TemplateRenderer struct {
	templates *template.Template
}

func main() {
	// create a channel to listen for signals
	sigChan := make(chan os.Signal, 1)

	// listen to SIGTERM and SIGINT (ctrl-c)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	templates, err := template.ParseFS(embeddedTemplates, "views/*.html")
	if err != nil {
		log.Fatalf("Error initializing templates: %v", err)
		os.Exit(1)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// serve server in a goroutine, allow the code to listen to ctrl-c
	go func() {
		e.Use(middleware.Logger())

		staticFiles, err := fs.Sub(embeddedFiles, "static")
		if err != nil {
			panic(err)
		}
		e.StaticFS("/static", staticFiles)

		e.GET("/", renderPage)

		e.GET("/service/health-check", func(c echo.Context) error {
			c.Response().Writer.WriteHeader(http.StatusOK)
			c.Response().Write([]byte("APP VERSION: 1.0.0"))
			return nil
		})

		port := os.Getenv("PORT")
		if port == "" {
			port = "5173"
		}

		e.Start(fmt.Sprintf(":%s", port))
	}()

	// block until a signal is received
	sig := <-sigChan
	log.Printf("Caught signal: %s\n", sig)

	// start graceful shutdown with a timeout
	// should stop everything and clean up within 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v\n", err)
	} else {
		log.Println("Server shut down gracefully.")
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
