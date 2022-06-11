package main

import (
	"github.com/gin-gonic/gin"
	"url-reducer/internal"
)

func main() {
	_, err := internal.InitDB()
	if err != nil {
		panic("Cannot connect to a database")
	}

	r := gin.Default()
	// TODO: mode from .env
	gin.SetMode(gin.DebugMode)
	r.GET("/api/read", internal.ReadUrl)
	r.POST("/api/put", internal.PutUrl)
	// TODO: port to .env
	r.Run(":8090")
}
