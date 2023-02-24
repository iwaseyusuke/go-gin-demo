package hello

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		name = "World"
	}
	c.Data(
		http.StatusOK,
		"text/plain",
		[]byte(fmt.Sprintf("Hello %s!", name)),
	)
}
