package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello": "Client",
	})
}
