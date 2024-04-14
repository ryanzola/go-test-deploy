package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/ryanzola/go-test-deploy/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading .env file", "err", err)
	}

	router := chi.NewMux()
	router.Get("/", handler.Make(handler.HandleHomeIndex))

	adminRouter := chi.NewMux()
	adminRouter.Get("/", handler.Make(handler.HandleAdminHomeIndex))

	router.Mount("/admin", adminRouter)

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Application is running on", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
