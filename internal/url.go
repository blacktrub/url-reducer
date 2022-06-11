package internal

type OriginUrl string
type ShortUrl string

// TODO: set limits
type URL struct {
	ID        uint `gorm:"primarykey"`
	OriginUrl OriginUrl
	ShortUrl  ShortUrl
}

func (url *URL) Create(originalUrl OriginUrl) URL {
}

func (u *URL) Get(shortUrl ShortUrl) (URL, error) {
	// TODO: create a connection somewhere else
	db, err := GetDBConnection()
	if err != nil {
		return URL{}, err
	}

	var url URL
	db.First(&url, "shortUrl = ?", shortUrl)
	return url, nil
}
