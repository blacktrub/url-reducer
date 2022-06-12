package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"url-reducer/internal/response"
)

func ReadUrl(c *gin.Context) {
	shortUrl := c.Query("hash")
	orgUrl, err := GetByShort(ShortUrl(shortUrl))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error("Not found"))
		return
	}
	c.JSON(http.StatusOK, response.Success(map[string]string{"url": string(orgUrl.OriginUrl)}))
}

func PutUrl(c *gin.Context) {
	// TODO: we need to decide how to check if url exists
	// TODO: maybe it's not a problem
	orgUrl := c.PostForm("url")
	if orgUrl == "" {
		c.JSON(http.StatusNotFound, response.Error("Bad request"))
		return
	}

	newUrl, err := CreateShort(OriginUrl(orgUrl))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error("Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, response.Success(map[string]string{"hash": string(newUrl.ShortUrl)}))
}
