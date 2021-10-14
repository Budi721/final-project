package router

import (
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/handler"
    "github.com/itp-backend/backend-a-co-create/service"
)

func NewRouter(dependencies service.Dependencies) *gin.Engine {
	router := gin.Default()

	setUserRouter(router, dependencies.UserService)
	return router
}

func setUserRouter(router *gin.Engine, dependencies service.IUserService) {
    router.POST("/login", handler.Login(dependencies))
    router.POST("/register", handler.Register(dependencies))
}

