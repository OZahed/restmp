package http

import (
	"log"

	"github.com/OZahed/restmp/internal/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	e *echo.Echo
}

func NewServer(l log.Logger) *Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.JWT(viper.GetString("jwt-key")))
	handlers.BindGreeting(e)

	return &Server{e: e}
}
