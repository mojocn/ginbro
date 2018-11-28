package handlers

import (
	"{{.OutPackage}}/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.POST("login", login)
}

func login(c *gin.Context) {
	var mdl models.{{.ModelName}}
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	ip := c.ClientIP()
	data, err := mdl.Login(ip)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
