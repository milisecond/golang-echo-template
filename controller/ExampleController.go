package controller

import "github.com/labstack/echo"

type ExampleController interface {
	Get(ctx echo.Context) error
	Post(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Find(ctx echo.Context) error
}

func NewExampleController() ExampleController {
	return ExampleControllerImpl{}
}
