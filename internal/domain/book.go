package domain

import "time"

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
