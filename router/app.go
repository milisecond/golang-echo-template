package router

import (
	"golang-echo-template/controller"

	"github.com/labstack/echo"
)

func App(router *echo.Echo) {
	exampleController := controller.NewExampleController()

	router.GET("/", exampleController.Get)
	router.POST("/", exampleController.Post)
	router.GET("/:user_id", exampleController.Find)
	router.DELETE("/:user_id", exampleController.Delete)
}
