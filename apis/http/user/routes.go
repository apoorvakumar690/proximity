package user

import (
	"proximity/shared"

	"github.com/gin-gonic/gin"
)

// NewUserRoute Creates and initializes user routes
func NewUserRoute(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := NewUserService(deps.Config, deps.Database, deps.HTTPRequester, deps.GrpcConn)
	userAPI := router.Group("/users")
	{
		userAPI.GET("/get-role", service.getAll)
		userAPI.GET("/get-access-list", service.getAll)
		userAPI.GET("/:userId", service.getOne)
	}
}
