package handlers

import (
	"desent/src/bootstrap"
	"desent/src/services"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	app *bootstrap.Application
	svc *services.Service
}

func NewHandler(app *bootstrap.Application) *Handler {
	return &Handler{
		app: app,
		svc: services.NewService(app),
	}
}

func errResp(c echo.Context, httpCode int, message string, errData any) error {
	return c.JSON(httpCode, map[string]any{
		"error": message,
		"data":  errData,
	})
}
