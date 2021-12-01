/*
Learn From :
https://www.youtube.com/watch?v=GjI0GSvmcSU&t=4884s
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Routing atau kalau Native GO ini adalah Handler
	router.GET("/", RootHandler)
	router.GET("/Product/:id", ProductHandler)       // ini Path Variable
	router.GET("/ProductQuery", ProductQueryHandler) // ini Query String

	router.Run() //bisa diubah port nya router.Run(":71"), secara default port nya adalah 8080
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Arief Kalbu",
		"nim":  "G64154030",
	})
}

func ProductHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func ProductQueryHandler(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
