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

func GetAllPosts(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	response, err := repository.GetPosts(ctx, client)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetAllRooms(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	response, err := repository.GetRooms(ctx, client)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetAllStatics(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	response, err := repository.GetStatics(ctx, client)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetAllTeams(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	response, err := repository.GetTeams(ctx, client)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetAllUsers(c *gin.Context) {

	ctx := context.Background()
	client, err := utils.CreateClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer client.Close()

	response, err := repository.GetUsers(ctx, client)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func AddPosts(c *gin.Context) {

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

	response, err := repository.PostPosts(ctx, client, body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func AddRooms(c *gin.Context) {

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

	response, err := repository.PostRooms(ctx, client, body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func AddStatic(c *gin.Context) {

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

	response, err := repository.PostStatics(ctx, client, body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func AddTeams(c *gin.Context) {

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

	response, err := repository.PostTeams(ctx, client, body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func AddUsers(c *gin.Context) {

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

	response, err := repository.PostUsers(ctx, client, body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdatePosts(c *gin.Context) {

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

	response, err := repository.PutPosts(ctx, client, c.Param("id"), body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateRooms(c *gin.Context) {

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

	response, err := repository.PutRooms(ctx, client, c.Param("id"), body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateStatic(c *gin.Context) {

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

	response, err := repository.PutStatics(ctx, client, c.Param("id"), body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateTeams(c *gin.Context) {

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

	response, err := repository.PutTeams(ctx, client, c.Param("id"), body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateUsers(c *gin.Context) {

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

	response, err := repository.PutUsers(ctx, client, c.Param("id"), body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
