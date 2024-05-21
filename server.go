package main

import (
	"fmt"
	"myapp/bingo"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	data := bingo.CreateGame()
	fmt.Println(data)

	jsonData, err := bingo.ConvertToJSON(data)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, string(jsonData))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
