package controller

import (
	"net/http"

	"golang-echo-template/model/request"
	"golang-echo-template/model/resource"

	"github.com/labstack/echo"
)

type ExampleControllerImpl struct{}

func (e ExampleControllerImpl) Get(ctx echo.Context) error {
	data := resource.User{
		Name:   "John Doe",
		UserId: "1",
		Age:    22,
	}

	return ctx.JSON(http.StatusOK, data)
}

func (e ExampleControllerImpl) Post(ctx echo.Context) error {
	var data request.User
	data = request.User{}

	err := ctx.Bind(&data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	return ctx.JSON(http.StatusOK, data)
}

// /user?user_id=11
func (e ExampleControllerImpl) Find(ctx echo.Context) error {
	user_id := ctx.QueryParam("user_id")

	return ctx.String(http.StatusOK, "user_id :"+user_id)
}

// /user_id
func (e ExampleControllerImpl) Delete(ctx echo.Context) error {
	user_id := ctx.Param("user_id")

	return ctx.String(http.StatusOK, "user_id :"+user_id)
}
