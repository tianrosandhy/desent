package handlers

import "github.com/labstack/echo/v4"

func (h *Handler) Index(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"OK": "Hello",
	})
}

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(200, map[string]any{
		"success": true,
	})
}

func (h *Handler) Echo(c echo.Context) error {
	return c.Stream(200, "application/json", c.Request().Body)
}
