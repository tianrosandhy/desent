package handlers

import (
	"desent/src/bootstrap"
	"desent/src/dto"
	"desent/src/pkg/response"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	req := dto.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse("Invalid request", err).Send(c, 400)
	}

	// hardcoded demo for now
	if req.Username != "admin" || req.Password != "password" {
		return response.ErrorResponse("Invalid credential").Send(c, 401)
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
		log.Printf("ERR Signed Token : %+v", err)
		return response.ErrorResponse("Cannot generate token").Send(c, 500)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": signedToken,
		"token":       signedToken,
	})
}
