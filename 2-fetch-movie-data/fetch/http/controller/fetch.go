package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stockbit/2-fetch-movie-data/fetch/http/request"
)

func GetAll(c *gin.Context) {
	search := c.Query("search")
	page := c.Query("page")

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", os.Getenv("OMDB_API_KEY"), search, page)
	result, err := request.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result.Map()})
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", os.Getenv("OMDB_API_KEY"), id)
	result, err := request.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result.Map()})
}
