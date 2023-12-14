package public_router

import (
	"collagen/controllers/api/v1/public_controller"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRouter(v1Router *gin.RouterGroup) {
	v1Router.GET("/app-info", public_controller.GetAppInfo)
}
