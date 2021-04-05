// Golang with Echo - PUT Method
// PUT 메소드는 기존 자원에 대한 변경(update)를 하는 행위이다.
// 기존 GET 메소드의 코드와 POST 메소드의 코드를 합쳐서 구현해 코드의 중복이 많다.
// 현재의 products는 map 형식이여서,

// PUT의 logic
// 1. products에서 해당하는 id의 product를 찾고 (GET)
// 2. 아이템을 변경한다. (유효성검사도 함께) (POST)

// TEST
// PUT http://localhost:8080/products/2 -> 원본의 값은 "tv"이다.
// {
// 	"product_name" : "speakers"
// }
// 변경후 GET을 해보면 변경된 값을 조회할 수 잇다.
// 만약, 초과된 번호 혹은 잘못된 id로 PUT을 하는 경우 -> 404 not found
// id가 유효석에 어긋난 경우 -> 500 internal server error

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

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
