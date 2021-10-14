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
    setArticleRouter(router, authMiddleware, dependencies.ArticleService)
	return router
}

func setUserRouter(router *gin.Engine, authMiddleware middleware.AuthValidate, dependencies service.IUserService) {
    router.POST("/login", authMiddleware.EnsureNotLoggedIn(), handler.Login(dependencies))
    router.POST("/register", authMiddleware.EnsureNotLoggedIn(), handler.Register(dependencies))
}

func setEnrollmentRouter(router *gin.Engine, authMiddleware middleware.AuthValidate, dependencies service.IEnrollmentService)  {
    enrollmentRouter := router.Group("/enrollment_requests")
    enrollmentRouter.GET("/", authMiddleware.EnsureLoggedIn(), handler.GetEnrollmentByStatus(dependencies))
    enrollmentRouter.POST("/approve", authMiddleware.EnsureLoggedIn(), handler.ApproveEnrollment(dependencies))
}

func setArticleRouter(router *gin.Engine, authMiddleware middleware.AuthValidate, dependencies service.IArticleService) {
    articleRouter := router.Group("/article")
    articleRouter.POST("/create", authMiddleware.EnsureLoggedIn(), handler.CreateArticleHandler(dependencies))
    articleRouter.GET("/detail/:id", authMiddleware.EnsureLoggedIn(), handler.GetArticleByIdHandler(dependencies))
    articleRouter.DELETE("/delete/:id", authMiddleware.EnsureLoggedIn(), handler.DeleteArticleHandler(dependencies))
    articleRouter.GET("/list_article", authMiddleware.EnsureLoggedIn(), handler.GetAllArticleHandler(dependencies))
}
