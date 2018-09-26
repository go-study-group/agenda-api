package main

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
)

type items struct {
	mongoSess mgo.Session
}

func (i items) list(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (i items) add(c echo.Context) error {
	return nil
}
