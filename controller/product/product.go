package ctrlProduct

import (
	"log"
	"net/http"

	"../../model"
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

	var product model.Product

	err := c.ShouldBindJSON(product)

	if err != nil {
		log.Fatal()
	}

	c.JSON(http.StatusOK, gin.H{
		"ProductName": product.ProductName,
	})
}
