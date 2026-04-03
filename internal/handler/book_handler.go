package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/theusualdeveloper/book-library/internal/application"
	"github.com/theusualdeveloper/book-library/internal/domain"
	"github.com/theusualdeveloper/book-library/internal/dto"
)

type BookHandler struct {
	bookService application.BookService
}

func NewBookHandler(bs application.BookService) BookHandler {
	return BookHandler{
		bookService: bs,
	}
}

func (bh BookHandler) GetListHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.List()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		bh.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(books); err != nil {
		bh.handleError(w, err)
		return
	}
}

func (bh BookHandler) FindHandler(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	book, err := bh.bookService.Find(uuid)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		bh.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(book); err != nil {
		bh.handleError(w, err)
		return
	}
}

func (bh BookHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}
	book, err := bh.bookService.Create(req)
	if err != nil {
		bh.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(book); err != nil {
		bh.handleError(w, err)
		return
	}
}

func (bh BookHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateRequest
	uuid := r.PathValue("id")
	err := json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}
	book, err := bh.bookService.Update(uuid, req)
	if err != nil {
		bh.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(book); err != nil {
		bh.handleError(w, err)
		return
	}
}

func (bh BookHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	err := bh.bookService.Delete(uuid)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		bh.handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (bh BookHandler) handleError(w http.ResponseWriter, err error) {
	var ve application.ValidationError
	if errors.As(err, &ve) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(ve.ToJSON())
		return
	}
	var de domain.DomainError
	if errors.As(err, &de) {
		switch de.Code {
		case domain.ErrCodeNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{"error": de.Message})
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
