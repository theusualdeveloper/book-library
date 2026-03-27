package port

import "github.com/theusualdeveloper/book-library/internal/domain"

type BookRepository interface {
	GetList() ([]domain.Book, error)
	Find(uuid string) (domain.Book, error)
	Create(b domain.Book) (domain.Book, error)
	Update(b domain.Book) (domain.Book, error)
	Delete(uuid string) error
}
