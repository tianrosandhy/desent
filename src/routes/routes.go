package routes

import (
	"desent/src/bootstrap"
	"desent/src/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, app *bootstrap.Application) {
	h := handlers.NewHandler(app)
	e.GET("/", h.Index)
	e.GET("/ping", h.Ping)
	e.POST("/echo", h.Echo)

	e.POST("/auth/token", h.Login)

	e.POST("/books", h.PostBook)
	e.GET("/books", h.GetAllBooks, tokenValidationMiddleware)
	e.GET("/books/:id", h.GetSingleBook)
	e.PUT("/books/:id", h.UpdateSingleBook)
	e.DELETE("/books/:id", h.DeleteSingleBook)

}
