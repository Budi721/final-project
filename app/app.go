package app

import (
	"github.com/itp-backend/backend-a-co-create/config"
)

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}

