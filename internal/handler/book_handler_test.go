package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theusualdeveloper/book-library/internal/application"
	"github.com/theusualdeveloper/book-library/internal/domain"
	"github.com/theusualdeveloper/book-library/internal/dto"
	"github.com/theusualdeveloper/book-library/internal/handler"
)

func TestGetListHandler(t *testing.T) {
	mock := &handler.MockService{
		BooksToReturn: []dto.BookResponse{
			{
				UUID:  "123",
				Title: "Go Programming",
			},
			{
				UUID:  "456",
				Title: "Clean Code",
			},
		},
	}

	h := handler.NewBookHandler(mock)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.GetListHandler(w, r)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}

	var books []dto.BookResponse
	json.NewDecoder(res.Body).Decode(&books)

	if len(books) != 2 {
		t.Fatalf("expected 2 books, got %d", len(books))
	}
}

func TestFindHandler(t *testing.T) {
	tests := []struct {
		name       string
		book       dto.BookResponse
		mockErr    error
		wantStatus int
	}{
		{
			name: "test 1",
			book: dto.BookResponse{
				UUID:  "123",
				Title: "Test Title",
			},
			mockErr:    nil,
			wantStatus: http.StatusOK,
		},
		{
			name: "test 2",
			book: dto.BookResponse{},
			mockErr: domain.DomainError{
				Code: domain.ErrCodeNotFound,
			},
			wantStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &handler.MockService{
				BookToReturn: tt.book,
				ErrToReturn:  tt.mockErr,
			}
			h := handler.NewBookHandler(mock)

			r := httptest.NewRequest(http.MethodGet, "/"+tt.book.UUID, nil)
			w := httptest.NewRecorder()
			h.FindHandler(w, r)

			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Fatalf("expected %d, got %d", tt.wantStatus, res.StatusCode)
			}
			if tt.mockErr == nil {
				var book dto.BookResponse
				json.NewDecoder(res.Body).Decode(&book)
				if book.UUID != mock.BookToReturn.UUID {
					t.Fatalf("expected book with uuid: %s, got: %s",
						mock.BookToReturn.UUID,
						book.UUID,
					)
				}
			}
		})
	}
}

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name       string
		req        dto.CreateRequest
		book       dto.BookResponse
		mockErr    error
		wantStatus int
	}{
		{
			name: "test 1",
			req: dto.CreateRequest{
				Title:         "Test Title",
				Author:        "Babak Shokouhi Pour",
				PublishedYear: 1989,
				Genre:         "Comedy",
				Pages:         500,
			},
			book: dto.BookResponse{
				UUID:  "123",
				Title: "Test Title",
			},
			mockErr:    nil,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "test 2",
			req:        dto.CreateRequest{},
			book:       dto.BookResponse{},
			mockErr:    application.ValidationError{},
			wantStatus: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &handler.MockService{
				BookToReturn: tt.book,
				ErrToReturn:  tt.mockErr,
			}

			h := handler.NewBookHandler(mock)
			b, _ := json.Marshal(tt.req)

			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			h.CreateHandler(w, r)
			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Fatalf("expected status code %d, got %d", tt.wantStatus, res.StatusCode)
			}
			if tt.mockErr == nil {
				var book dto.BookResponse
				json.NewDecoder(res.Body).Decode(&book)
				if book.UUID != mock.BookToReturn.UUID {
					t.Fatalf("expected book with uuid: %s, got: %s",
						mock.BookToReturn.UUID,
						book.UUID,
					)
				}
			}
		})
	}
}

func TestDeleteHandler(t *testing.T) {
	tests := []struct {
		name       string
		book       dto.BookResponse
		wantStatus int
		mockErr    error
	}{
		{
			name: "test 1",
			book: dto.BookResponse{
				UUID:  "123",
				Title: "Test Title",
			},
			wantStatus: http.StatusNoContent,
			mockErr:    nil,
		},
		{
			name:       "test 2",
			book:       dto.BookResponse{},
			wantStatus: http.StatusNotFound,
			mockErr: domain.DomainError{
				Code: domain.ErrCodeNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &handler.MockService{
				BookToReturn: tt.book,
				ErrToReturn:  tt.mockErr,
			}

			h := handler.NewBookHandler(mock)

			r := httptest.NewRequest(http.MethodDelete, "/"+tt.book.UUID, nil)
			w := httptest.NewRecorder()

			h.DeleteHandler(w, r)

			res := w.Result()

			if res.StatusCode != tt.wantStatus {
				t.Fatalf("expected status %d got %d", tt.wantStatus, res.StatusCode)
			}
		})
	}

}
