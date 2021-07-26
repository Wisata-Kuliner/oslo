package handlers

import (
	"context"
	"net/http"

	"github.com/Wisata-Kuliner/oslo/internal/repository"
	utils "github.com/Wisata-Kuliner/oslo/internal/utils"
	"github.com/gin-gonic/gin"
)

// func Test(w http.ResponseWriter, r *http.Request) {
func Test(c *gin.Context) {
	ctx := context.Background()
	app, err := utils.NewConfig(ctx)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	if err != nil {
		// fmt.Fprintf(w, err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	client, err := app.Firestore(ctx)
	// fmt.Printf("BELOM %v", (client == nil))
	if err != nil {
		// log.Fatalln(err)
		// c.String(http.StatusInternalServerError, err.Error())
		// return
		client, err = utils.CreateClient(ctx)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	defer client.Close()

	response, err := repository.GetPlayers(ctx, client)
	if err != nil {
		// fmt.Fprintf(w, err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
