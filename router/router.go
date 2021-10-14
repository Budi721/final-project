package router

import (
    "github.com/gin-gonic/gin"
    "github.com/itp-backend/backend-a-co-create/handler"
    "github.com/itp-backend/backend-a-co-create/middleware"
    "github.com/itp-backend/backend-a-co-create/service"
)

func NewRouter(dependencies service.Dependencies) *gin.Engine {
	router := gin.Default()
    authMiddleware := dependencies.AuthValidate
	setUserRouter(router, authMiddleware, dependencies.UserService)
    setEnrollmentRouter(router, authMiddleware, dependencies.EnrollmentService)
	return router
}

func setUserRouter(router *gin.Engine, authMiddleware middleware.AuthValidate, dependencies service.IUserService) {
    router.POST("/login", authMiddleware.EnsureNotLoggedIn(), handler.Login(dependencies))
    router.POST("/register", authMiddleware.EnsureNotLoggedIn(), handler.Register(dependencies))
}

func setEnrollmentRouter(router *gin.Engine, authMiddleware middleware.AuthValidate, dependencies service.IEnrollmentService)  {
    router.GET("/requests", authMiddleware.EnsureLoggedIn(), handler.GetEnrollmentByStatus(dependencies))
    router.POST("/approve", authMiddleware.EnsureLoggedIn(), handler.ApproveEnrollment(dependencies))
}
