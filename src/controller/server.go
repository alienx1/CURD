package controller

import (
	"api/src/config/server/cors"
	td "api/src/controller/td"

	"github.com/gin-gonic/gin"
)

func ServerRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.AllowCORS())
	api := r.Group("api")
	{
		td.Td(api)
	}
	return r
}
