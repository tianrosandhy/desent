package services

import "desent/src/bootstrap"

type Service struct {
	app *bootstrap.Application
}

func NewService(app *bootstrap.Application) *Service {
	return &Service{
		app: app,
	}
}
