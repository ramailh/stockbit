package route

import (
	"github.com/gin-gonic/gin"
	"github.com/stockbit/2-fetch-movie-data/fetch/http/controller"
	"github.com/stockbit/2-fetch-movie-data/fetch/http/middleware"
)

func NewRouter() *gin.Engine {
	rtr := gin.Default()

	rtr.Use(gin.LoggerWithFormatter(middleware.LoggerMiddleware))

	rtr.GET("/", controller.GetAll)
	rtr.GET("/:id", controller.GetByID)

	return rtr
}
