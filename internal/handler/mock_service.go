package handler

import (
	"github.com/theusualdeveloper/book-library/internal/dto"
)

type MockService struct {
	BooksToReturn []dto.BookResponse
	BookToReturn  dto.BookResponse
	ErrToReturn   error
}

func (m *MockService) List() ([]dto.BookResponse, error) {
	return m.BooksToReturn, m.ErrToReturn
}

func (m *MockService) Find(uuid string) (dto.BookResponse, error) {
	return m.BookToReturn, m.ErrToReturn
}

func (m *MockService) Create(req dto.CreateRequest) (dto.BookResponse, error) {
	return m.BookToReturn, m.ErrToReturn
}

func (m *MockService) Update(uuid string, req dto.UpdateRequest) (dto.BookResponse, error) {
	return m.BookToReturn, m.ErrToReturn
}

func (m *MockService) Delete(uuid string) error {
	return m.ErrToReturn
}
