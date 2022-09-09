package http

import (
	"context"
	"fmt"

	ilog "github.com/OZahed/restmp/internal/log"

	"github.com/OZahed/restmp/internal/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
	}))
	handlers.BindGreeting(e)

	return &Server{e: e}
}

func (s *Server) ListenAndServe(ctx context.Context, addr string) error {
	s.e.Server.Addr = addr
	ilog.Logger.Debug(fmt.Sprintf("listening on : %s", addr))
	return s.e.Server.ListenAndServe()
}
