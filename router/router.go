package router

import (
	"gynocentric/handler"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.Engine {
	api := engine.Group("api")
	{
		api.GET("/", handler.Home)
	}

	return engine
}
