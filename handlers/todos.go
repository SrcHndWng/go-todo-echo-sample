package handlers

import (
	"net/http"
	"strconv"

	"github.com/SrcHndWng/go-todo-echo-sample/repository"
	"github.com/labstack/echo"
)

// CreateTodo Handler
func CreateTodo(c echo.Context) error {
	t := repository.NewTodo()
	if err := c.Bind(t); err != nil {
		return err
	}
	repository.AddTodo(t)
	return c.JSON(http.StatusCreated, t)
}

// GetTodos Handler
func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, repository.GetTodos())
}

// GetTodo Handler
func GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	t := repository.GetTodo(id)
	if t == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, t)
}

// UpdateTodo Handler
func UpdateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	t := new(repository.Todo)
	if err := c.Bind(t); err != nil {
		return err
	}
	update := repository.UpdateTodo(id, t)
	if update == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, update)
}

// DeleteTodo Handler
func DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	todos := repository.DeleteTodo(id)
	if todos == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, todos)
}
