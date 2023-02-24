package main

import (
	"embed"
	"html/template"

	"github.com/iwaseyusuke/go-gin-demo/pkg/hello"

	"github.com/gin-gonic/gin"
)

//go:embed web/templates/*.html
var f embed.FS

func main() {
	r := gin.Default()
	t := template.Must(template.New("").ParseFS(f, "web/templates/*.html"))
	r.SetHTMLTemplate(t)
	r.GET("/hello", hello.Hello)
	r.Run()
}
