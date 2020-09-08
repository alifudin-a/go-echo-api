package routers

import (
	handler "github.com/alifudin-a/go-echo-api/handler/character"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init for inititalize router
func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")

	v1.GET("/characters", handler.List)

	return e
}
