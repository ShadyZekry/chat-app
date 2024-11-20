package middlewares

import (
	"net/http"

	"github.com/ShadyZekry/chat-app/services"
	"github.com/labstack/echo/v4"
)

func ValidateApplication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Param("token")
			if !services.ValidateToken(token) {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
			}

			return next(c)
		}
	}
}
