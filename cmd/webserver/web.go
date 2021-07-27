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

	router.GET("/players", handlers.Test)
	router.GET("/posts", handlers.GetAllPosts)
	router.GET("/rooms", handlers.GetAllRooms)
	router.GET("/static", handlers.GetAllStatics)
	router.GET("/teams", handlers.GetAllTeams)
	router.GET("/users", handlers.GetAllUsers)

	return router
}
