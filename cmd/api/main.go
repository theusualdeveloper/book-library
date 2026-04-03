package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/theusualdeveloper/book-library/internal/adapter"
	"github.com/theusualdeveloper/book-library/internal/application"
	"github.com/theusualdeveloper/book-library/internal/handler"
)

func main() {
	repository := adapter.NewInMemoryRepository()
	service := application.NewBookService(repository)
	bookHandler := handler.NewBookHandler(service)

	books := http.NewServeMux()
	books.HandleFunc("GET /", bookHandler.GetListHandler)
	books.HandleFunc("GET /{id}", bookHandler.FindHandler)
	books.HandleFunc("POST /", bookHandler.CreateHandler)
	books.HandleFunc("PUT /{id}", bookHandler.UpdateHandler)
	books.HandleFunc("DELETE /{id}", bookHandler.DeleteHandler)

	mux := http.NewServeMux()
	mux.Handle("/books/", http.StripPrefix("/books", books))

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	fmt.Println("Server is starting on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
