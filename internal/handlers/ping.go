package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
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

func Prepare(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	response, err := repository.PostPlayers(ctx, client, body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func Execute(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	response, err := repository.PutPlayers(ctx, client, c.Param("id"), body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func Download(c *gin.Context) {
	b, err := utils.ReadServiceAccount()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	response, err := repository.DownloadData(b)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", response)
}
