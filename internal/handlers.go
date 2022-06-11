package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadUrl(c *gin.Context) {
	shortUrl := c.Param("hash")
	url := URL{}
	orgUrl, err := url.Get(ShortUrl(shortUrl))
	if err != nil {
		c.String(http.StatusNotFound, "not found")
		return
	}
	c.String(http.StatusOK, string(orgUrl.OriginUrl))
}
