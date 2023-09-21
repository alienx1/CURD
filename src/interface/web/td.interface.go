package web

import (
	"api/src/config/server/send.server"
	"api/src/model"
	"api/src/service/td_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TdInterface interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
type Td struct{}

var err error

func (td *Td) Get(c *gin.Context) {
	var u []model.UserModel
	if err = td_service.GetService(&u); err != nil {
		send.SendModel(c, "fail", err, http.StatusBadRequest)

	} else {
		send.SendModel(c, &u, nil, http.StatusOK)
	}
}
func (td *Td) Delete(c *gin.Context) {
	user := c.Param("user")
	if err = td_service.Delete(user); err != nil {
		send.SendModel(c, "fail", err, http.StatusBadRequest)
	} else {
		send.SendModel(c, "pass", nil, http.StatusOK)

	}
}

func (td *Td) Post(c *gin.Context) {
	var u model.UserModel
	if err = c.ShouldBindJSON(&u); err != nil {
		send.SendModel(c, "fail", err, http.StatusBadRequest)
	} else {
		if err = td_service.PostService(&u); err != nil {
			send.SendModel(c, "fail", err, http.StatusBadRequest)
		} else {
			send.SendModel(c, "pass", nil, http.StatusOK)

		}
	}
}

func (td *Td) Put(c *gin.Context) {
	var u model.UserModel
	if err = c.ShouldBindJSON(&u); err != nil {
		send.SendModel(c, "fail", err, http.StatusBadRequest)
	} else {
		if err = td_service.PutService(&u); err != nil {
			send.SendModel(c, "fail", err, http.StatusBadRequest)
		} else {

			send.SendModel(c, "pass", nil, http.StatusOK)
		}
	}
}

func NewTdInterface() TdInterface {
	return &Td{}
}
