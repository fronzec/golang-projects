package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Hello World!"))
	})

	r.Mount("/posts", postsResource{}.Routes())
	log.Fatal(http.ListenAndServe(":"+port, r))
}
