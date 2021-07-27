package handlers

import (
	"context"
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
