package routes

import (
	"Shakirsadiq6/CRUD-APIs-Go/auth"
	"Shakirsadiq6/CRUD-APIs-Go/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoute(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route to generate token
	e.POST("/login", auth.Login)

	// Restricted group
	restricted := e.Group("/user")
	{
		restricted.Use(middleware.JWT([]byte("secret")))
		restricted.POST("", (&controllers.Handler{}).CreateUser)
		restricted.GET("/:userId", (&controllers.Handler{}).GetAUser)
		restricted.PUT("/:userId", (&controllers.Handler{}).EditAUser)
		restricted.DELETE("/:userId", (&controllers.Handler{}).DeleteAUser)
		restricted.GET("/all", (&controllers.Handler{}).GetAllUsers)
	}

}
