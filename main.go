package main

import (
	"net/http"
	"personal-site/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "github.com/labstack/echo-jwt/v4"
)

func main() {
  e := echo.New()

  e.Use(middleware.CORS())
  e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
  e.Use(middleware.BodyLimit("4M"))
  e.Use(echojwt.JWT([]byte("secret")))

  adminGroup := e.Group("/admin")

  e.GET("/", func (c echo.Context) error {
    return c.String(http.StatusOK, "Hello World")
  })

  adminGroup.POST("/", controllers.SignIn)
  adminGroup.POST("/", controllers.SignUp)

  e.Logger.Fatal(e.Start(":3001"))
}
