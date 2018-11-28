package handlers

import (
	"github.com/dejavuzhou/ginbro/boilerplate/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("department", departmentAll)
	groupApi.GET("department/:id", departmentOne)
	groupApi.POST("department", departmentCreate)
	groupApi.PATCH("department", departmentUpdate)
	groupApi.DELETE("department/:id", departmentDelete)
}

func departmentAll(c *gin.Context) {
	mdl := models.Department{}
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
func departmentOne(c *gin.Context) {
	var mdl models.Department
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
func departmentCreate(c *gin.Context) {
	var mdl models.Department
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

func departmentUpdate(c *gin.Context) {
	var mdl models.Department
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

func departmentDelete(c *gin.Context) {
	var mdl models.Department
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
