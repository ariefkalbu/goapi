/*
Learn From :
https://www.youtube.com/watch?v=GjI0GSvmcSU&t=4884s
*/
package main

import (
	"github.com/gin-gonic/gin"

	ctrlRoot "goapi/controller"
	ctrlProduct "goapi/controller/product"
)

func main() {
	router := gin.Default()

	//Routing atau kalau Native GO ini adalah Handler
	router.GET("/", ctrlRoot.RootHandler)
	router.GET("/Product/:id/:name", ctrlProduct.ProductHandler) // ini Path Variable
	router.GET("/ProductQuery", ctrlProduct.ProductQueryHandler) // ini Query String
	router.GET("/ProductPost/", ctrlProduct.ProductPostHandler)  // ini POST

	router.Run() //bisa diubah port nya router.Run(":71"), secara default port nya adalah 8080
}
