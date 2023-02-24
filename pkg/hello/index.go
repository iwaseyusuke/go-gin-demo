package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		name = "World"
	}
	c.HTML(http.StatusOK, "hello.html", gin.H{
		"name": name,
	})
}
