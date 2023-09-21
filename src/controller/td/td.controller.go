package td_controller

import (
	"api/src/interface/web"

	"github.com/gin-gonic/gin"
)

var T = web.NewTdInterface()

func Td(api *gin.RouterGroup) {
	td := api.Group("/curd")
	{
		td.GET("/get", T.Get)
		td.POST("/post", T.Post)
		td.PUT("/put", T.Post)
		td.DELETE("/delete/:user", T.Delete)
	}
}
