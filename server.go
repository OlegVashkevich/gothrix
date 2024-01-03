package main

import (
	//"net/http"
	"gothrix/components/component"

	"github.com/OlegVashkevich/templ_components/element"
	"github.com/labstack/echo/v4"
)

// test
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return component.Hello("John").Render(c.Request().Context(), c.Response())
		//return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return element.H1(id).Render(c.Request().Context(), c.Response())
		//return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
