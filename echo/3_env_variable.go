// Golang with Echo - env variable
// 환경변수를 이용해서 port를 설정하는 등 코드의 유연함을 위해 사용된다.
// 그리고, 컨테이너 기반으로 동작할 때, 환경변수의 사용이 많다.
// default 값도 설정할 수 있다.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("MY_APP_PORT") // MY_APP_PORT라는 변수가 있으면 값을 가져온다.
	if port == "" {                  // default는 8080
		port = "8080"
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Well, hello there !!")
	})
	e.Logger.Print(fmt.Sprintf("Listening on port %s", port)) // 각 port의 변경 값도 변수로 처리한다.
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
