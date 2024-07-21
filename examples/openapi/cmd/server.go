package main

import (
	"github.com/canidam/examples/openapi/internal/api"
	"github.com/canidam/examples/openapi/internal/users"
	"github.com/labstack/echo/v4"
)

type Server struct {
	users.UsersHandler
}

func main() {
	e := echo.New()
	s := Server{UsersHandler: *users.New(nil)}

	api.RegisterHandlers(e, &s)
	e.Logger.Fatal(e.Start(":8080"))
}
