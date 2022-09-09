package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.StaticFile("/css", "./templates")

	router.Use(static.Serve("/", static.LocalFile("./templates", false)))

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
