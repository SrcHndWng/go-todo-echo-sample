package main

import (
	"github.com/SrcHndWng/go-todo-echo-sample/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.File("/", "templates/index.html")
	e.POST("/todos", handlers.CreateTodo)
	e.GET("/todos", handlers.GetTodos)
	e.GET("/todos/:id", handlers.GetTodo)
	e.PUT("/todos/:id", handlers.UpdateTodo)
	e.DELETE("/todos/:id", handlers.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
