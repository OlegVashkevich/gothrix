package main

import (
	//"net/http"
	"gothrix/components/common"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return common.Hello("John").Render(c.Request().Context(), c.Response())
		//return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return common.Hello(id).Render(c.Request().Context(), c.Response())
		//return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
