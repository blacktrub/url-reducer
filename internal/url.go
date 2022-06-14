package internal

import (
	"encoding/base32"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type OriginUrl string
type ShortUrl string

// TODO: set limits
type URL struct {
	ID        uint `gorm:"primarykey"`
	OriginUrl OriginUrl
	ShortUrl  ShortUrl `gorm:"index"`
}

func createHash(id uint) ShortUrl {
	hash := base32.StdEncoding.EncodeToString([]byte(fmt.Sprint(id)))
	return ShortUrl(strings.ReplaceAll(hash, "=", ""))
}

func CreateShort(db *gorm.DB, originalUrl OriginUrl) (URL, error) {
	url := URL{OriginUrl: originalUrl, ShortUrl: ""}
	db.Create(&url)
	url.ShortUrl = createHash(url.ID)
	db.Save(&url)
	return url, nil
}

func GetByShort(db *gorm.DB, shortUrl ShortUrl) (URL, error) {
	var url URL
	result := db.Take(&url, "short_url = ?", shortUrl)
	if result.Error != nil {
		return url, result.Error
	}
	return url, nil
}
