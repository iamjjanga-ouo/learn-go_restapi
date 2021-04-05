// Golang with Echo - route parameters
// URI에 resource에 대한 접근을 parameter를 통해서 할 수 있다.
// 그리고 이 parameter에 대한 roting역시 할 수 있어 마음껏 처리가 가능하다.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	// products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well, hello there!!")
	})
	e.GET("/products", func(c echo.Context) error {
		// return c.JSON(http.StatusOK, products)
		return c.JSON(http.StatusOK, []string{"apple", "banana", "mango"})
	})
	e.GET("/products/:vendor", func(c echo.Context) error {
		// return c.JSON(http.StatusOK, c.Param("vendor")) //  [parameter에 대한 routing] GET http://localhost:8080/products/apple -> "apple"
		return c.JSON(http.StatusOK, c.QueryParam("olderThan")) // [query에 대한 routing] GET http://localhost:8080/products/apple?olderThan=iphone10 -> "iphone10"
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
