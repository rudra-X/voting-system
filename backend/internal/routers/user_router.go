package routers

import (
	"voiting-system/internal/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", services.GetUsers)
	r.POST("/register", services.RegisterUser)
	r.POST("/login", services.Login)
}
