package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1Apis := r.Group("/v1")
	UserRoutes(v1Apis.Group("/user"))
	return r
}
