package application_test

import (
	"errors"
	"testing"

	"github.com/theusualdeveloper/book-library/internal/adapter"
	"github.com/theusualdeveloper/book-library/internal/application"
	"github.com/theusualdeveloper/book-library/internal/domain"
	"github.com/theusualdeveloper/book-library/internal/dto"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name    string
		req     dto.CreateRequest
		wantErr bool
	}{
		{
			name: "test 1",
			req: dto.CreateRequest{
				Title:         "How to kill Khamene'ei",
				Author:        "Donald Trump",
				PublishedYear: 2026,
				Genre:         "Reality",
				Pages:         1,
			},
			wantErr: false,
		},
		{
			name:    "test 2",
			req:     dto.CreateRequest{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &adapter.MockRepository{}
			service := application.NewBookService(repo)
			res, err := service.Create(tt.req)
			if !tt.wantErr {
				if err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
				if res.Title != tt.req.Title {
					t.Errorf("expected title %s, got %s", tt.req.Title, res.Title)
				}
				if res.UUID == "" {
					t.Error("expected UUID to be set")
				}
			} else {
				if err == nil {
					t.Fatal("expected validation error, got nil")
				}
				var ve application.ValidationError
				if !errors.As(err, &ve) {
					t.Fatalf("expected ValidationError, got %T", err)
				}
				if len(ve.Errs) == 0 {
					t.Error("expected validation errors to be present")
				}
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		book    domain.Book
		mockErr error
		wantErr bool
	}{
		{
			name:    "test 1",
			book:    domain.NewBook("How to success", "Donald trump", "Comedy", 120, 2023),
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "test 2",
			book: domain.Book{},
			mockErr: domain.DomainError{
				Code:    domain.ErrCodeNotFound,
				Message: "book not found",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &adapter.MockRepository{
				BookToReturn: tt.book,
				ErrToReturn:  tt.mockErr,
			}
			service := application.NewBookService(repo)
			res, err := service.Find(tt.book.UUID)
			if !tt.wantErr {
				if err != nil {
					t.Fatalf("expected no error but got: %v", err)
				}
				if res.UUID != tt.book.UUID {
					t.Fatalf("want book with uuid: %s, got: %s", tt.book.UUID, res.UUID)
				}
				if res.UUID == "" {
					t.Error("expected UUID to be set")
				}
			} else {
				if err == nil {
					t.Fatalf("expected domain error got nil")
				}
				var de domain.DomainError
				if !errors.As(err, &de) {
					t.Fatalf("expected domain error got different error type")
				}
				var expectedDe domain.DomainError
				if errors.As(tt.mockErr, &expectedDe) {
					if de.Code != expectedDe.Code {
						t.Fatalf("expected error code %s, got %s", expectedDe.Code, de.Code)
					}
					if de.Message != expectedDe.Message {
						t.Fatalf("expected error message: %s, got: %s", expectedDe.Message, de.Message)
					}
				}
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name      string
		req       dto.UpdateRequest
		book      domain.Book
		mockError error
	}{
		{
			name: "test 1",
			req: dto.UpdateRequest{
				Title:         "Test 1 editted",
				Author:        "Babak editted",
				PublishedYear: 2025,
				Genre:         "Comedy/Criminal",
				Pages:         455,
			},
			book:      domain.NewBook("Test 1", "Babak", "Comedy", 1989, 349),
			mockError: nil,
		},
		{
			name:      "test 2",
			req:       dto.UpdateRequest{},
			book:      domain.NewBook("Test 2", "Bashir", "Drama", 1982, 434),
			mockError: application.ValidationError{},
		},
		{
			name: "test 3",
			req: dto.UpdateRequest{
				Title:         "Test 2",
				Author:        "Bashir",
				PublishedYear: 1989,
				Genre:         "Drama",
				Pages:         123,
			},
			book: domain.Book{},
			mockError: domain.DomainError{
				Code: domain.ErrCodeNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &adapter.MockRepository{
				ErrToReturn: tt.mockError,
			}
			service := application.NewBookService(repo)
			_, err := service.Update(tt.book.UUID, tt.req)
			switch expectedDe := tt.mockError.(type) {
			case nil:
				if err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
			case domain.DomainError:
				var de domain.DomainError
				if !errors.As(err, &de) {
					t.Fatalf("expected domain error got %v", err)
				}
				if de.Code != expectedDe.Code {
					t.Fatalf("expected not found err got %s", de.Code)
				}
			case application.ValidationError:
				var ve application.ValidationError
				if !errors.As(err, &ve) {
					t.Fatalf("expected validation error got %v", err)
				}
			default:
				t.Fatalf("invalid mock error type")
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		book      domain.Book
		mockError error
	}{
		{
			name:      "test 1",
			book:      domain.NewBook("Test 1", "Babak", "Comedy", 1989, 349),
			mockError: nil,
		},
		{
			name: "test 2",
			book: domain.Book{},
			mockError: domain.DomainError{
				Code: domain.ErrCodeNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &adapter.MockRepository{
				ErrToReturn: tt.mockError,
			}
			service := application.NewBookService(repo)
			err := service.Delete(tt.book.UUID)
			switch expectedErr := tt.mockError.(type) {
			case nil:
				if err != nil {
					t.Fatalf("expected no error got: %v", err)
				}
			case domain.DomainError:
				var de domain.DomainError
				if !errors.As(err, &de) {
					t.Fatalf("expected domain error got: %v", err)
				}
				if de.Code != expectedErr.Code {
					t.Fatalf("expected not found error got %s", de.Code)
				}
			default:
				t.Fatalf("invalid specified mock error")
			}
		})
	}
}
