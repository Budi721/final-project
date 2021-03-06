package service

import (
	"github.com/itp-backend/backend-a-co-create/app"
	"github.com/itp-backend/backend-a-co-create/external/jwt_client"
	"github.com/itp-backend/backend-a-co-create/middleware"
	"github.com/itp-backend/backend-a-co-create/repository"
)

type Dependencies struct {
	AuthValidate      middleware.AuthValidate
	UserService       IUserService
	EnrollmentService IEnrollmentService
	ArticleService    IArticleService
	ProjectService    IProjectService
}

func InstantiateDependencies(application *app.Application) Dependencies {
	jwtWrapper := jwt_client.New()

	authMiddleware := middleware.NewAuthValidate(application.Config, jwtWrapper)
	userRepo := repository.NewUserRepository(application.DB)
	userService := NewUserService(userRepo, application.Config, jwtWrapper)
	enrollmentRepo := repository.NewEnrollmentRepository(application.DB)
	enrollmentService := NewEnrollmentService(enrollmentRepo)
	articleRepo := repository.NewArticleRepository(application.DB)
	articleService := NewArticleService(articleRepo)
	projectRepo := repository.NewProjectRepository(application.DB)
	projectService := NewProjectService(projectRepo)

	return Dependencies{
		AuthValidate:      authMiddleware,
		UserService:       userService,
		EnrollmentService: enrollmentService,
		ArticleService:    articleService,
		ProjectService:    projectService,
	}
}
