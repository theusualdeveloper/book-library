package application

import (
	"fmt"

	"github.com/theusualdeveloper/book-library/internal/domain"
	"github.com/theusualdeveloper/book-library/internal/dto"
	"github.com/theusualdeveloper/book-library/internal/port"
)

type BookService struct {
	bookRepository port.BookRepository
}

func NewBookService(repository port.BookRepository) BookService {
	return BookService{
		bookRepository: repository,
	}
}

func (bs BookService) List() ([]dto.BookResponse, error) {
	books, err := bs.bookRepository.GetList()
	if err != nil {
		return nil, fmt.Errorf("get list of books failed: %w", err)
	}
	return mapEntitiesToResponse(books), nil
}

func (bs BookService) Find(uuid string) (dto.BookResponse, error) {
	book, err := bs.bookRepository.Find(uuid)
	if err != nil {
		return dto.BookResponse{}, fmt.Errorf("find book failed: %w", err)
	}
	return mapEntityToResponse(book), nil
}

func (bs BookService) Create(req dto.CreateRequest) (dto.BookResponse, error) {
	errs := req.Validate()
	if len(errs) > 0 {
		return dto.BookResponse{}, ValidationError{Errs: errs}
	}
	book := domain.NewBook(req.Title, req.Author, req.Genre, req.Pages, req.PublishedYear)
	book, err := bs.bookRepository.Create(book)
	if err != nil {
		return dto.BookResponse{}, err
	}
	return mapEntityToResponse(book), nil
}

func (bs BookService) Update(uuid string, req dto.UpdateRequest) (dto.BookResponse, error) {
	errs := req.Validate()
	if len(errs) > 0 {
		return dto.BookResponse{}, ValidationError{Errs: errs}
	}
	b, err := bs.bookRepository.Find(uuid)
	if err != nil {
		return dto.BookResponse{}, err
	}
	b = b.Update(req.Title, req.Author, req.Genre, req.Pages, req.PublishedYear)
	b, err = bs.bookRepository.Update(b)
	if err != nil {
		return dto.BookResponse{}, err
	}
	return mapEntityToResponse(b), nil
}

func (bs BookService) Delete(uuid string) error {
	return bs.bookRepository.Delete(uuid)
}

func mapEntitiesToResponse(books []domain.Book) []dto.BookResponse {
	booksRes := []dto.BookResponse{}
	for _, b := range books {
		booksRes = append(booksRes, dto.BookResponse{
			UUID:          b.UUID,
			Title:         b.Title,
			Author:        b.Author,
			Genre:         b.Genre,
			PublishedYear: b.PublishedYear,
			Pages:         b.Pages,
			CreatedAt:     b.CreatedAt,
			UpdatedAt:     b.UpdatedAt,
		})
	}
	return booksRes
}

func mapEntityToResponse(book domain.Book) dto.BookResponse {
	return dto.BookResponse{
		UUID:          book.UUID,
		Title:         book.Title,
		Author:        book.Author,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Pages:         book.Pages,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}
}
