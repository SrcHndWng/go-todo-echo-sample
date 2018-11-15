package main

import (
	"github.com/SrcHndWng/go-todo-echo-sample/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/todos", handlers.CreateTodo)
	e.GET("/todos/:id", handlers.GetTodo)
	e.Logger.Fatal(e.Start(":1323"))
}