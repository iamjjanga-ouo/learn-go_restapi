// Golang with Echo - starting Echo
// echo
// - echo는 엄청 큰 struct이다.
// - e := echo 여기서 e는 echo구조체를 가리키는 pointer

// e.GET(path, handleFunc, ...)
// - HandleFunc(context) 현재 request의 context이다. -> interface

// e.Logger.Print() / e.Logger.Fatal()
// Logger defines the logging interface.

// GET:localhost:8080/ -> 200(OK)와 Well, Hello there~! 출력
// GET:localhost:8080/hello -> 404(Not Found)
// POST:localhost:8080/ -> 405(Method not allowed)

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well, Hello there~!") // http.StatusOK는 200을 의미
	})
	// err := e.Start(":8080")
	e.Logger.Print("Listening on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
