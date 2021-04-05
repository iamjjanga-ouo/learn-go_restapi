// Golang with Echo - Validator
// go get gopkg.in/go-playground/validator.v9

// REST API 설계에서 data validation은 request data에 대한 유효성 검사이다.
// 즉, 특정 request(JSON)의 value에 어떠한 정보가 정해진 규칙에 맞게 지정되어야 메소드를 실행할 수 있다는 의미.
// 이 코드의 예시에서는 validator.v9라는 외부 모듈을 통해서 진행했고,
// 규칙에 관련된것은 공식 문서를 따라 작성하면된다.

// POST JSON example
// {
// 	"product_name" : "tablet",
// 	"vendor" : "apple",
// 	"email" : "support@apple.com",
// 	"website" : "http://apple.com",
// 	"country" : "US",
// 	"default_device_ip" : "192.168.8.8"
// }

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

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

		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
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
			// Vendor          string `json:"vendor" validate:"min=5,max=10"`
			// Email           string `json:"email" validate:"required_with=Vendor,email"`
			// Website         string `json:"website" validate:"url"`
			// Country         string `json:"country" validate:"len=2"`
			// DefaultDeviceIp string `json:"default_device_ip" validate:"ip"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}

		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
