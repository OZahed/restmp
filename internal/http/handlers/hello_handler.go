package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	gh     = &GreetingsHandler{}
	prefix = "api"
)

type GreetingsHandler struct{}

func BindGreeting(e *echo.Echo) {
	g := e.Group(prefix)
	g.GET("/", gh.Greet)
}

func (g *GreetingsHandler) Greet(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Greetings form base server",
	})
}
