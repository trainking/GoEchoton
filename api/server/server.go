package server

import (
	"GoEchoton/api/server/middleware"
	"GoEchoton/api/server/router"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e    *echo.Echo
	mFun []echo.MiddlewareFunc
}

func (s *Server) Start() {
	for _, f := range s.mFun {
		s.e.Use(f)
	}
	router.Initized(s.e)
	s.e.Logger.Fatal(s.e.Start(":1323"))
}

func New() *Server {
	_e := echo.New()
	server := &Server{
		e:    _e,
		mFun: middleware.Initized(),
	}
	return server
}
