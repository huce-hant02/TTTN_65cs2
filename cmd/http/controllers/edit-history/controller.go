package edit_history

import "github.com/gin-gonic/gin"

type Controller interface {
	GetInfo(ctx *gin.Context)
}

// RegisterRoutesV1 register routes for version 1
func RegisterRoutesV1(router *gin.Engine, controller Controller) {
	v1 := router.Group("/v1")
	v1.GET("/get-info", controller.GetInfo)
}
