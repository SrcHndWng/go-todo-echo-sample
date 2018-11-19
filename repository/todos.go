package repository

import (
	"strconv"

	"github.com/labstack/echo"
)

type (
	Todo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	todos = make([]Todo, 0)
	seq   = 1
)

func AddTodo(c echo.Context) (*Todo, error) {
	t := &Todo{
		ID: seq,
	}
	if err := c.Bind(t); err != nil {
		return nil, err
	}
	todos = append(todos, *t)
	seq++
	return t, nil
}

func GetTodos() []Todo {
	return todos
}

func GetTodo(c echo.Context) (*Todo, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}
	for _, t := range todos {
		if id == t.ID {
			return &t, nil
		}
	}
	return nil, nil
}
