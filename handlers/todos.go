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
	todos = map[int]*todo{}
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
	todos[t.ID] = t
	seq++
	return c.JSON(http.StatusCreated, t)
}

// GetTodo Handler
func GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, todos[id])
}
