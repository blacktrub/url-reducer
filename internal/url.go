package internal

import (
	"encoding/base32"
	"fmt"
	"strings"
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

func createHash(id uint) ShortUrl {
	hash := base32.StdEncoding.EncodeToString([]byte(fmt.Sprint(id)))
	return ShortUrl(strings.ReplaceAll(hash, "=", ""))
}

func CreateShort(originalUrl OriginUrl) (URL, error) {
	// TODO: create a connection somewhere else
	db, err := GetDBConnection()
	if err != nil {
		return URL{}, err
	}

	url := URL{OriginUrl: originalUrl, ShortUrl: ""}
	db.Create(&url)
	url.ShortUrl = createHash(url.ID)
	db.Save(&url)
	return url, nil
}

func GetByShort(shortUrl ShortUrl) (URL, error) {
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
