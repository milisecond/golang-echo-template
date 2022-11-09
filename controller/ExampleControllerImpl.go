package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type ExampleControllerImpl struct{}

func (e ExampleControllerImpl) Get(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Get")
}

func (e ExampleControllerImpl) Post(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Post")

}

func (e ExampleControllerImpl) Find(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Find")

}

func (e ExampleControllerImpl) Delete(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Delete")

}
