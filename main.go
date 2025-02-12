package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"snippetbox.aragorn.net/handlers"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	listenAddr := os.Getenv("LISTEN_ADDRESS")
	slog.Info("HTTP Server is running", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)

}
