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
	t := selectTodo(id)
	return t, nil
}

func UpdateTodo(c echo.Context) (*Todo, error) {
	t := new(Todo)
	if err := c.Bind(t); err != nil {
		return nil, err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	update := selectTodo(id)
	update.Name = t.Name
	return update, nil
}

func selectTodo(id int) *Todo {
	for _, t := range todos {
		if id == t.ID {
			return &t
		}
	}
	return nil
}
