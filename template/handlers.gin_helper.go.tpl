package handlers

import (
	"errors"
	"{{.OutPackage}}/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func jsonError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(200, gin.H{"code": 0, "msg": msg})
}
func jsonData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"code": 1, "data": data})
}
func jsonPagination(c *gin.Context, list interface{}, total uint, query *models.PaginationQuery) {
	c.JSON(200, gin.H{"code": 1, "data": list, "total": total, "offset": query.Offset, "limit": query.Limit})
}
func jsonSuccess(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "msg": "success"})
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func parseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseId), nil
}

func enableCorsMiddleware() {
	//TODO:: customize your own CORS
	//https://github.com/gin-contrib/cors
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //https://foo.com
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, //enable cookie
		AllowOriginFunc: func(origin string) bool {
			return true
			//return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour, //cache options result decrease request lag
	}))
}
