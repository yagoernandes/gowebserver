package main

import (
	"gowebserver/src/api/albums"
	"gowebserver/src/api/home"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", home.HelloHandler)
	router.GET("/albums", albums.GetAlbumsHandler)

	router.Run("0.0.0.0:8080")
}
