package handlers

import (
	"github.com/dejavuzhou/ginbro/boilerplate/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("user", jwtMiddleware, userAll)
	groupApi.GET("user/:id", jwtMiddleware, userOne)
	groupApi.POST("user", jwtMiddleware, userCreate)
	groupApi.PATCH("user", jwtMiddleware, userUpdate)
	groupApi.DELETE("user/:id", jwtMiddleware, userDelete)
}

func userAll(c *gin.Context) {
	mdl := models.AuthorizationModel{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}
func userOne(c *gin.Context) {
	var mdl models.AuthorizationModel
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
func userCreate(c *gin.Context) {
	var mdl models.AuthorizationModel
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}

	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

func userUpdate(c *gin.Context) {
	var mdl models.AuthorizationModel
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

func userDelete(c *gin.Context) {
	var mdl models.AuthorizationModel
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
