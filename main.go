package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Result struct {
	Second string `json:"second"`
	Minute string `json:"minute"`
}

func main() {
	e := echo.New()
	e.GET("/api/v1/get_date_calculator/:startDate/:endDate", func(c echo.Context) error {

		var result Result

		return c.JSON(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
