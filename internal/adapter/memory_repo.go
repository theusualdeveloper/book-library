package adapter

import (
	"fmt"

	"github.com/theusualdeveloper/book-library/internal/domain"
)

type InMemoryRepository struct {
	books map[string]domain.Book
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		books: map[string]domain.Book{},
	}
}

func (im *InMemoryRepository) GetList() ([]domain.Book, error) {
	books := []domain.Book{}
	for _, b := range im.books {
		books = append(books, b)
	}
	return books, nil
}

func (im *InMemoryRepository) Find(uuid string) (domain.Book, error) {
	book, ok := im.books[uuid]
	if !ok {
		return domain.Book{}, fmt.Errorf("book not found: %s", uuid)
	}
	return book, nil
}

func (im *InMemoryRepository) Create(b domain.Book) (domain.Book, error) {
	im.books[b.UUID] = b
	return b, nil
}

func (im *InMemoryRepository) Update(b domain.Book) (domain.Book, error) {
	_, ok := im.books[b.UUID]
	if !ok {
		return domain.Book{}, fmt.Errorf("book not found: %s", b.UUID)
	}
	im.books[b.UUID] = b
	return b, nil
}

func (im *InMemoryRepository) Delete(uuid string) error {
	_, ok := im.books[uuid]
	if !ok {
		return fmt.Errorf("book not found: %s", uuid)
	}
	delete(im.books, uuid)
	return nil
}
