package routes

import (
	"io"

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
		rawBody, _ := io.ReadAll(c.Request().Body)
		return c.String(200, string(rawBody))
	})
}
