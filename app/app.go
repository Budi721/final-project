package app

import (
    "github.com/itp-backend/backend-a-co-create/config"
    "gorm.io/gorm"
)

type Application struct {
	Config *config.Config
    DB *gorm.DB
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
        DB: config.InitDB(),
	}

	return application
}

