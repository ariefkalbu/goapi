package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Routing atau kalau Native GO ini adalah Handler
	router.GET("/", RootHandler)

	router.Run() //bisa diubah port nya router.Run(":71"), secara default port nya adalah 8080
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Arief Kalbu",
		"nim":  "G64154030",
	})
}
