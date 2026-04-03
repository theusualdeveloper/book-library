package dto

import (
	"time"
)

type CreateRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"year"`
	Genre         string `json:"genre"`
	Pages         int    `json:"pages"`
}

func (cr CreateRequest) Validate() map[string][]string {
	return validateFields(cr.Title, cr.Author, cr.Genre, cr.Pages, cr.PublishedYear)
}

type UpdateRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"year"`
	Pages         int    `json:"pages"`
}

func (ur UpdateRequest) Validate() map[string][]string {
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

func validateFields(title, author, genre string, pages, year int) map[string][]string {
	errs := map[string][]string{}
	if title == "" {
		errs["title"] = append(errs["title"], "title is required")
	}
	if author == "" {
		errs["author"] = append(errs["author"], "author is required")
	}
	if year == 0 {
		errs["year"] = append(errs["year"], "published year is required")
	}
	if genre == "" {
		errs["genre"] = append(errs["genre"], "genre is required")
	}
	if pages == 0 {
		errs["pages"] = append(errs["pages"], "pages is required")
	}
	return errs
}
