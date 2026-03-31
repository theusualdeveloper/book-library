package dto

import (
	"errors"
	"time"
)

type CreateRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"year"`
	Genre         string `json:"genre"`
	Pages         int    `json:"pages"`
}

func (cr CreateRequest) Validate() error {
	return validateFields(cr.Title, cr.Author, cr.Genre, cr.Pages, cr.PublishedYear)
}

type UpdateRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"year"`
	Pages         int    `json:"pages"`
}

func (ur UpdateRequest) Validate() error {
	return validateFields(ur.Title, ur.Author, ur.Genre, ur.Pages, ur.PublishedYear)
}

type BookResponse struct {
	UUID          string    `json:"uuid"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Genre         string    `json:"genre"`
	PublishedYear int       `json:"year"`
	Pages         int       `json:"pages"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func validateFields(title, author, genre string, pages, year int) error {
	var errs []error
	if title == "" {
		errs = append(errs, errors.New("title is required"))
	}
	if author == "" {
		errs = append(errs, errors.New("author is required"))
	}
	if year == 0 {
		errs = append(errs, errors.New("published year is required"))
	}
	if genre == "" {
		errs = append(errs, errors.New("genre is required"))
	}
	if pages == 0 {
		errs = append(errs, errors.New("pages is required"))
	}
	return errors.Join(errs...)
}
