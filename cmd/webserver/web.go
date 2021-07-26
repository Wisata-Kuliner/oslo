package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handlers "github.com/Wisata-Kuliner/oslo/internal/handlers"
)

func NewRouter(port string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/test", handlers.Test)

	return router
}
