package api

import (
	"example/echoserver/database"
	"example/echoserver/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	body := new(types.CreateUser)
	if err := c.Bind(body); err != nil {
		return err
	}

	newUser := types.User{
		Id:       "3",
		UserName: body.UserName,
		Email:    body.Email,
	}
	database.STORE = append(database.STORE, newUser)
	return c.JSON(http.StatusCreated, newUser)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, database.STORE)
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
