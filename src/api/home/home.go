package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct{}{})
}
