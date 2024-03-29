package main

import (
	"github.com/iwaseyusuke/go-gin-demo/pkg/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/hello", hello.Hello)
	r.Run()
}
