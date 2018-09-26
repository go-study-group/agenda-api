package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	conf, err := unmarshalConfig()
	if err != nil {
		log.Fatalf("config error (%s)", err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	items := &items{}

	apiGroup := e.Group("/api")
	apiGroup.GET("/:group", items.list)
	apiGroup.POST("/:group", items.add)

	// Start server
	hostString := fmt.Sprintf(":%s", conf.Port)
	e.Logger.Fatal(e.Start(hostString))
}
