package main

import (
	"os"
	"url-reducer/internal"

	"github.com/gin-gonic/gin"
)

func getEnv(key string, d string) string {
	val := os.Getenv(key)
	if val == "" {
		return d
	}
	return val
}

func main() {
	db, err := internal.SetupDB()
	if err != nil {
		panic("Cannot connect to a database")
	}

	h := internal.GetHandler(db)

	r := gin.Default()
	r.SetTrustedProxies([]string{getEnv("APP_HOST", "127.0.0.1")})
	gin.SetMode(getEnv("APP_MODE", gin.DebugMode))
	r.GET("/api/read", h.ReadUrl)
	r.POST("/api/put", h.PutUrl)
	r.Run(getEnv("APP_PORT", ":8090"))
}
