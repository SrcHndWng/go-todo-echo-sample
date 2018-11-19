package handlers

import (
	"net/http"

	"github.com/SrcHndWng/go-todo-echo-sample/repository"
	"github.com/labstack/echo"
)

// CreateTodo Handler
func CreateTodo(c echo.Context) error {
	t, err := repository.AddTodo(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, t)
}

// GetTodos Handler
func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, repository.GetTodos())
}

// GetTodo Handler
func GetTodo(c echo.Context) error {
	t, err := repository.GetTodo(c)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, nil)
	}
	if t == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, t)
}

// UpdateTodo Handler
func UpdateTodo(c echo.Context) error {
	t, err := repository.UpdateTodo(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, t)
}
