package service

import (
	"github.com/itp-backend/backend-a-co-create/app"
	"github.com/itp-backend/backend-a-co-create/external/jwt_client"
	"github.com/itp-backend/backend-a-co-create/external/mysql"
	"github.com/itp-backend/backend-a-co-create/middleware"
	"github.com/itp-backend/backend-a-co-create/repository"
)

type Dependencies struct {
	AuthValidate middleware.AuthValidate
	UserService  IUserService
}

func InstantiateDependencies(application *app.Application) Dependencies {
	jwtWrapper := jwt_client.New()
	mysqlClient := mysql.NewMysqlClient(mysql.ClientConfig{
		Username: application.Config.DBUsername,
		Password: application.Config.DBPassword,
		Host:     application.Config.DBHost,
		Port:     application.Config.DBPort,
		DBName:   application.Config.DBName,
	})
	db := mysqlClient.OpenDB()
	authMiddleware := middleware.NewAuthValidate(application.Config, jwtWrapper)
	userRepo := repository.NewUserRepository(db)
	userService := NewUserService(userRepo, application.Config, jwtWrapper)
	return Dependencies{
        AuthValidate: authMiddleware,
		UserService: userService,
	}
}
