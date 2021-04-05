//HandleFunc
// [!] HandlerFunc와 HandleFunc는 이름이 비슷하지만 다른함수이다.

// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// - localhost:8080/hello와 localhost:8080/login 그리고 각 endpoint에 대해서 GET, POST, DELETE 등 각각 다르게 동작해야한다.
// - 하드웨어 상이나, 다르게 Multiplexer를 이용하지만, echo Framework를 사용하기전 간단하게 HandleFunc로 구현해본다.
package main

import (
	"fmt"
	"net/http"
)

func myLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
		<head>
			Login
		</head>
		<body>
			<h1>Please enter your username and password</h1>
		</body>
	</html>
	`)
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
		<head>
			Hi
		</head>
		<body>
			<h1>Welcome</h1>
		</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/login", myLogin)
	// http.HandleFunc("/welcome", myWelcome) // welcome 뒤에 /가 붙는지 않붙는지에 따라 차이가 있다.
	http.HandleFunc("/welcome/", myWelcome)   // "/welcome" -> /welcome/123 접근이 불가하다 (404)
	fmt.Println("Listening on port 8080....") // "/welcome/" -> /welcome/123 접근이 가능하다.
	http.ListenAndServe("localhost:8080", nil)
}

// installed gin : go get github.com/codegangsta/gin
