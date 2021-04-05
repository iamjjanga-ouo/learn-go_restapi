//Handle
// - Handle과 HandleFunc의 차이점은 2번째 parameter로 어떤 타입을 취급하냐이다.
// func Handle(pattern string, handler Handler)
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
package main

import (
	"fmt"
	"net/http"
)

type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "on login page")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "on welcome page")
}

// func myLogin(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
// 	<!DOCTYPE html>
// 	<html>
// 		<head>
// 			Login
// 		</head>
// 		<body>
// 			<h1>Please enter your username and password</h1>
// 		</body>
// 	</html>
// 	`)
// }

// func myWelcome(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
// 	<!DOCTYPE html>
// 	<html>
// 		<head>
// 			Hi
// 		</head>
// 		<body>
// 			<h1>Welcome</h1>
// 		</body>
// 	</html>
// 	`)
// }

func main() {
	// http.HandleFunc("/login", myLogin)
	// http.HandleFunc("/welcome/", myWelcome)
	// http.Handle("/login", http.HandlerFunc(myLogin))     // HandlerFunc(f) == handler
	// http.Handle("/welcome", http.HandlerFunc(myWelcome)) // HandlerFunc도 handler interface가 가지는 ServeHTTP method를 내장하고 있다.
	var i login
	var j welcome
	http.Handle("/login", i)
	http.Handle("/welcome", j)
	fmt.Println("Listening on port 8080....")
	http.ListenAndServe("localhost:8080", nil)
}

// installed gin : go get github.com/codegangsta/gin
