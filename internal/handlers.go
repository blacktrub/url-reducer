package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadUrl(c *gin.Context) {
	shortUrl := c.Query("hash")
	orgUrl, err := GetByShort(ShortUrl(shortUrl))
	if err != nil {
		c.String(http.StatusNotFound, "not found")
		return
	}
	c.String(http.StatusOK, string(orgUrl.OriginUrl))
}

func PutUrl(c *gin.Context) {
	// TODO: we need to decide how to check if url exists
	// TODO: maybe it's not a problem
	orgUrl := c.PostForm("url")
	if orgUrl == "" {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	newUrl, err := CreateShort(OriginUrl(orgUrl))
	if err != nil {
		c.String(http.StatusInternalServerError, "smth went wrong")
		return
	}
	c.String(http.StatusOK, string(newUrl.ShortUrl))
}
