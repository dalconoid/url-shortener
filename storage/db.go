package storage

import (
	"fmt"
	"github.com/dalconoid/url-shortener/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

//Database represents a real db
type Database struct {
	Db *gorm.DB
	ConnString string
}

func (db *Database) Open() error {
	var err error
	db.Db, err = gorm.Open(postgres.Open(db.ConnString), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) RegisterURL(url string, slug string) (*models.ShortenedURL, *models.CustomError) {
	su := &models.ShortenedURL{URL: url, Slug: slug}

	result := db.Db.FirstOrCreate(&su, models.RegisterURLRequest{URL: url})
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), models.SlugNotUniqueMsg) {
			return nil, &models.CustomError{Err: fmt.Errorf("slug [%s] already binded", slug), Code: models.ErrSlugNotUnique}
		} else {
			return nil, &models.CustomError{Err: result.Error, Code: models.ErrDefaultCode}
		}
	}

	return su, nil
}

func (db *Database) GetRedirectURL(slug string) (string, *models.CustomError) {
	redirectURL := ""

	su := &models.ShortenedURL{}
	result := db.Db.Where(models.ShortenedURL{Slug: slug}).First(su)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return redirectURL, &models.CustomError{Err: fmt.Errorf("slug [%v] not registered", slug), Code: models.ErrSlugNotRegistered}
		}
		return redirectURL, &models.CustomError{Err: result.Error, Code: models.ErrDefaultCode}
	}
	redirectURL = su.URL

	return redirectURL, nil
}

func (db *Database) getShortenedUrl(url string) (*models.ShortenedURL, *models.CustomError) {
	su := &models.ShortenedURL{}
	db.Db.Where(&models.ShortenedURL{URL: url})
	return su, nil
}