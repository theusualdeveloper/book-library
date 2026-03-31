package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	UUID          string
	Title         string
	Author        string
	PublishedYear int
	Genre         string
	Pages         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewBook(title, author, genre string, pages, year int) Book {
	return Book{
		UUID:          uuid.NewString(),
		Title:         title,
		Author:        author,
		PublishedYear: year,
		Genre:         genre,
		Pages:         pages,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (b Book) Update(title, author, genre string, pages, year int) Book {
	b.Title = title
	b.Author = author
	b.Genre = genre
	b.PublishedYear = year
	b.Pages = pages
	b.UpdatedAt = time.Now()
	return b
}
