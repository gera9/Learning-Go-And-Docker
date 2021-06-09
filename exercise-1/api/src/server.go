package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := "3000"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Server on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Home"))
	})

	r.Mount("/people", PeopleResource{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))
}
