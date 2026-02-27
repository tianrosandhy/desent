package routes

import (
	"desent/src/bootstrap"
	"desent/src/dto"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func tokenValidationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(401, map[string]string{
				"error": "missing Authorization header",
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(401, map[string]string{
				"error": "invalid token format",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &dto.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return bootstrap.JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(401, map[string]string{
				"error": "invalid or expired token",
			})
		}

		claims := token.Claims.(*dto.JwtCustomClaims)
		c.Set("username", claims.Username)

		return next(c)
	}
}
