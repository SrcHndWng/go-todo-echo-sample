package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	todo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	todos = make([]todo, 0)
	seq   = 1
)

// CreateTodo Handler
func CreateTodo(c echo.Context) error {
	t := &todo{
		ID: seq,
	}
	if err := c.Bind(t); err != nil {
		return err
	}
	todos = append(todos, *t)
	seq++
	return c.JSON(http.StatusCreated, t)
}

// GetTodos Handler
func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

// GetTodo Handler
func GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, t := range todos {
		if id == t.ID {
			return c.JSON(http.StatusOK, t)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}
