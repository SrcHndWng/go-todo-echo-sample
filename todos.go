package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getTodo(c echo.Context) error {
	id := c.Param("id")
	result := fmt.Sprintf("todo id = %s", id)
	return c.String(http.StatusOK, result)
}
