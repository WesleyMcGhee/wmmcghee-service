package main

import (
	"net/http"
	"personal-site/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  e := echo.New()

  e.Use(middleware.CORS())

  e.GET("/", func (c echo.Context) error {
    return c.String(http.StatusOK, "Hello World")
  })

  e.POST("/auth/signup", controllers.SignUp)
  e.POST("/auth/signin", controllers.SignIn)

  e.Logger.Fatal(e.Start(":3001"))
}
