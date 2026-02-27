package bootstrap

import "gorm.io/gorm"

type Application struct {
	DB *gorm.DB
}

func NewApplication() *Application {
	app := Application{
		DB: SetupDB(),
	}

	return &app
}
