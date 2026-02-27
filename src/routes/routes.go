package routes

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
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
}
