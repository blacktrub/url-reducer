package main

import (
	"url-reducer/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := internal.SetupDB()
	if err != nil {
		panic("Cannot connect to a database")
	}

	h := internal.GetHandler(db)

	r := gin.Default()
	// TODO: mode from .env
	gin.SetMode(gin.DebugMode)
	r.GET("/api/read", h.ReadUrl)
	r.POST("/api/put", h.PutUrl)
	// TODO: port to .env
	r.Run(":8090")
}
