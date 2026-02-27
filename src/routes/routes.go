package routes

import (
	"desent/src/bootstrap"
	"desent/src/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, app *bootstrap.Application) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"OK": "Hello",
		})
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"success": true,
		})
	})

	e.POST("/echo", func(c echo.Context) error {
		return c.Stream(200, "application/json", c.Request().Body)
	})

	h := handlers.NewHandler(app)
	e.POST("/books", h.PostBook)
	e.GET("/books", h.GetAllBooks)
	e.GET("/books/:id", h.GetSingleBook)
	e.PUT("/books/:id", h.UpdateSingleBook)
	e.DELETE("/books/:id", h.DeleteSingleBook)

}
