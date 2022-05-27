package main

import (
	"Shakirsadiq6/CRUD-APIs-Go/configs"
	"Shakirsadiq6/CRUD-APIs-Go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":6000"))
}
