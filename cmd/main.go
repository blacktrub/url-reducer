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
	r.GET("/api/read", internal.ReadUrl)
	// TODO: port to .env
	r.Run(":8090")
}
