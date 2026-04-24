package adapter

import (
	"github.com/theusualdeveloper/book-library/internal/domain"
)

type MockRepository struct {
	BookToReturn  domain.Book
	BooksToReturn []domain.Book
	ErrToReturn   error
}

func (m *MockRepository) GetList() ([]domain.Book, error) {
	return m.BooksToReturn, m.ErrToReturn
}

func (m *MockRepository) Find(uuid string) (domain.Book, error) {
	return m.BookToReturn, m.ErrToReturn
}

func (m *MockRepository) Create(b domain.Book) (domain.Book, error) {
	return b, m.ErrToReturn
}

func (m *MockRepository) Update(b domain.Book) (domain.Book, error) {
	return b, m.ErrToReturn
}

func (m *MockRepository) Delete(uuid string) error {
	return m.ErrToReturn
}
