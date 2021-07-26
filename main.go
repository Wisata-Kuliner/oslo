package main

import (
	// "log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"

	web "github.com/Wisata-Kuliner/oslo/cmd/webserver"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "8080"
	}

	router := web.NewRouter(port)

	router.Run(":" + port)
}
