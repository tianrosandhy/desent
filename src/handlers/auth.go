package handlers

import (
	"desent/src/bootstrap"
	"desent/src/dto"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	req := dto.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return errResp(c, 400, "Invalid request", err)
	}

	// hardcoded demo for now
	if req.Username != "admin" || req.Password != "password" {
		return errResp(c, 401, "Invalid credentials", nil)
	}

	claims := &dto.JwtCustomClaims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(bootstrap.JWT_SECRET)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not generate token",
		})
	}

	// test case doesnt give any response format, so return as string for now
	return c.String(http.StatusOK, signedToken)
}
