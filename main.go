package main

import (
	"golang-echo-template/conf"
	"golang-echo-template/router"

	"github.com/labstack/echo"
)

var r *echo.Echo

func init() {
	r = echo.New()

	router.App(r)
}

func main() {
	r.Start(":" + conf.PORT)
}
