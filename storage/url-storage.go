package storage

import "github.com/dalconoid/url-shortener/models"

type IURLStorage interface {
	RegisterURL(string, string) (*models.ShortenedURL, *models.CustomError)
	GetRedirectURL(string) (string, *models.CustomError)
}
