package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(port string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	return router
}