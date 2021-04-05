// Golang with Echo - DELETE method
// DELETE 메서드를 이용해 특정 원소를 삭제할 수 있다.
// GET 메서드와 유사한 코드형식을 가지나, 특정 index를 얻고 그 인덱스를 map형식에서 지우는 방법도 추가되어야한다.

// Example
// DELETE  http://localhost:8080/products/2 -> "tv"
// 후 GET을 통한 조회
// [
//   {
//     "1": "mobiles"
//   },
//   {
//     "3": "laptops"
//   }
// ]

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// ProductValidator echo validator for product
type ProductValidator struct {
	validator *validator.Validate
}

// Validate validates product request body
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	v := validator.New()
	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well, hello there!!")
	})
	// GET /products는 전체 resource를 return해준다.
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})
	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string

		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for _, p := range products {
			for k := range p {

				if pID == k {
					product = p
				}
			}
		}
		// Not matched items
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}

		var reqBody body
		e.Validator = &ProductValidator{validator: v}

		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}

		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})

	e.PUT("/products/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for _, p := range products {
			for k := range p {
				if pID == k {
					product = p
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "product not fount")
		}

		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var reqBody body
		e.Validator = &ProductValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product[pID] = reqBody.Name
		return c.JSON(http.StatusOK, product)
	})

	e.DELETE("/products/:id", func(c echo.Context) error {
		var product map[int]string
		var index int

		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		for i, p := range products {
			for k := range p {

				if pID == k {
					product = p
					index = i
				}
			}
		}
		// Not matched items
		if product == nil {
			return c.JSON(http.StatusNotFound, "product not found")
		}

		// Splice map : delete element in map
		splice := func(s []map[int]string, index int) []map[int]string {
			return append(s[:index], s[index+1:]...) // idiomatic in go
		}
		products = splice(products, index)

		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
