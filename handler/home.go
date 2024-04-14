package handler

import (
	"log/slog"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Welcome to the home page! Cool!"))

	return nil
}

func HandleAdminHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// log the request path
	slog.Info("Request path", "path", r.URL.Path)
	w.Write([]byte("Welcome to the admin home page! Cool!"))

	return nil
}
