package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwt := c.Request().Header.Get("x-auth-token")
		if jwt == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorized",
			})
		}

		// --@@-- check database for the user if it is accurate
		c.Set("User", jwt)
		return next(c)
	}
}
