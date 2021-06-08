package models

import "time"

const (
	ErrDefaultCode = 0
	ErrURLNotUnique = 1
	ErrSlugNotUnique = 2
	ErrSlugNotRegistered = 3

	URLNotUniqueMsg = "unique_url"
	SlugNotUniqueMsg = "unique_slug"
)

type CustomError struct {
	Err error
	Code int
}

type ShortenedURL struct {
	ID int `gorm:"column:id"`
	Slug string
	URL string
	CreatedAt time.Time
}

func (ShortenedURL) TableName() string {
	return "urls"
}

type RegisterURLRequest struct {
	URL string `json:"url" validate:"required"`
	Slug string `json:"slug"`
}

type RegisterURLResponse struct {
	Slug string `json:"slug"`
}