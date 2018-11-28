package handlers

import (
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("{{.ResourceName}}",{{.HandlerName}}All)
	groupApi.GET("{{.ResourceName}}/:id", {{.HandlerName}}One)
	groupApi.POST("{{.ResourceName}}", {{.HandlerName}}Create)
	groupApi.PATCH("{{.ResourceName}}", {{.HandlerName}}Update)
	groupApi.DELETE("{{.ResourceName}}/:id", {{.HandlerName}}Delete)
}
//All
func {{.HandlerName}}All(c *gin.Context) {

}
//One
func {{.HandlerName}}One(c *gin.Context) {


}
//Create
func {{.HandlerName}}Create(c *gin.Context) {

}
//Update
func {{.HandlerName}}Update(c *gin.Context) {

}
//Delete
func {{.HandlerName}}Delete(c *gin.Context) {

}
