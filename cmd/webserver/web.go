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

	router.GET("/storage", handlers.Download)

	router.POST("/players", handlers.Prepare)
	router.POST("/posts", handlers.AddPosts)
	router.POST("/rooms", handlers.AddRooms)
	router.POST("/static", handlers.AddStatic)
	router.POST("/teams", handlers.AddTeams)
	router.POST("/users", handlers.AddUsers)

	router.PUT("/players/:id", handlers.Execute)
	router.PUT("/posts/:id", handlers.UpdatePosts)
	router.PUT("/rooms/:id", handlers.UpdateRooms)
	router.PUT("/static/:id", handlers.UpdateStatic)
	router.PUT("/teams/:id", handlers.UpdateTeams)
	router.PUT("/users/:id", handlers.UpdateUsers)

	return router
}
