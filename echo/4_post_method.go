// Golang with Echo - POST Method
// POST Method를 사용해 resource를 추가하는 예시

// OUTPUT
// POST localhost:8080/products + JSON { "product_name" : "headphone" }
// GET localhost:8080/products
// [
//   {
//     "1": "mobiles"
//   },
//   {
//     "2": "tv"
//   },
//   {
//     "3": "laptops"
//   },
//   {
//     "4": "headphone"
//   }
// ]

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

	// POST /products는 새로운 resource product를 추가한다.
	e.POST("/products", func(c echo.Context) error {
		// body 구조체를 생성, JSON 타입의 Name 필드를 가진다.
		type body struct {
			Name string `json:"product_name"`
		}
		var reqBody body
		// 요청 본문을 Go type으로 바인딩 하려면 다음을 사용해야한다. Context#Bind(i interface{})
		// 디폴트 바인더는 Content-Type 해더를 기반으로 application/json, application/xml 및 application/x-www-form-urlencoded 데이터의 디코딩을 지원한다
		// 아래의 예제는 요청 페이러도를 body 구조체에 바인딩한다.
		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		product := map[int]string{ // 기존 collection의 길이에 1을 더한 값을 key로하고,
			len(products) + 1: reqBody.Name, // 바인딩했던 body타입의 Name을 value로 하는 dictionary 객체를 생성한다.
		}
		products = append(products, product)  // 기존의 products collection에 새로운 document를 추가한다.
		return c.JSON(http.StatusOK, product) // 추가한 document를 출력한다.
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
