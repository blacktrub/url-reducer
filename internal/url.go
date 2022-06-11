package internal

import (
	"encoding/base32"
)

type OriginUrl string
type ShortUrl string

// TODO: set limits
type URL struct {
	ID        uint `gorm:"primarykey"`
	OriginUrl OriginUrl
	// TODO: add an index because we will query by the field
	ShortUrl ShortUrl
}

func (u *URL) createHash(id uint) ShortUrl {
	return ShortUrl(base32.StdEncoding.EncodeToString([]byte(string(id))))
}

func (u *URL) Create(originalUrl OriginUrl) (URL, error) {
	// TODO: create a connection somewhere else
	db, err := GetDBConnection()
	if err != nil {
		return URL{}, err
	}

	url := URL{OriginUrl: originalUrl, ShortUrl: ""}
	db.Create(&url)
	url.ShortUrl = u.createHash(url.ID)
	db.Save(&url)
	return url, nil
}

func (u *URL) Get(shortUrl ShortUrl) (URL, error) {
	// TODO: create a connection somewhere else
	db, err := GetDBConnection()
	if err != nil {
		return URL{}, err
	}

	var url URL
	result := db.Take(&url, "short_url = ?", shortUrl)
	if result.Error != nil {
		return url, result.Error
	}
	return url, nil
}
