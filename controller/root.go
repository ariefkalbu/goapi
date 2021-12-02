package ctrlRoot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Arief Kalbu",
		"nim":  "G64154030",
	})
}
