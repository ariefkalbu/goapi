package ctrlProduct

import (
	"goapi/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func ProductPostHandler(c *gin.Context) {
	var product model.ProductModel

	err := c.ShouldBindJSON(&product)

	if err != nil {
		log.Fatal()
	}

	c.JSON(http.StatusOK, gin.H{
		"ProductName": product.ProductName,
		"Price":       product.Price,
		"Description": product.ProductDescription,
	})
}
