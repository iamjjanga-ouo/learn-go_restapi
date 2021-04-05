// Golang with Echo - expanding GET
// GET 메서드를 사용하여 실제 REST API에 요청 예시를 구현한다.

// OUTPUT
// localhost:8080/products/1 -> { "1": "mobiles" }
// localhost:8080/products/2 -> { "2": "tv" }
// localhost:8080/products/3 -> { "3": "laptops" }
// localhost:8080/products/4 -> product not found

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}} // colletion : products, document: mobiles, tv, labtops

	e.GET("/", func(c echo.Context) error { // Get endpoint root path -> check connectino health (Status OK)
		return c.String(http.StatusOK, "Well, hello there!!")
	})
	e.GET("/products/:id", func(c echo.Context) error { // Get resource using URI
		var product map[int]string

		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id")) // Convert string("id") to int("id")
				if err != nil {
					return err
				}
				if pID == k { // Match wanted product's id
					product = p // then store object in product
				}
			}
		}
		// Not matched items
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusOK, product)
	})
	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
