//http methods net/http
// - endpoint로 접근을 위한 다양한 방법 (GET, POST, ...)를 가진 request를 구분하는 방법
// `r.Method`를 이용하여 처리한다.
// switch ~ case문을 이용하면 더 깔끔한 처리가 가능하다. (하지만, 복잡성에서는 선호하지 않는 방법이다.)
// Production 환경에서나, 실무 API Server를 구현할 때, switch문을 사용하면 complexity한 처리를 하지 못할 수도있다.

// 또 다른 방법으로는 different endpoint를 처리하는 Function을 두는 방법이 현명하다.
package main

import (
	"fmt"
	"net/http"
)

func myLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "Using GET for login endpoint")
	}
	if r.Method == "POST" {
		fmt.Fprintln(w, "Using POST for login endpoint")
	}
	fmt.Fprint(w, "on login page")
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	// switch문을 이용
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Using GET for welcome endpoint.")
	case "POST":
		fmt.Fprintln(w, "Using POST for welcome endpoint.")
	}
	fmt.Fprint(w, "on welcome page")
}

func main() {
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome/", myWelcome)
	fmt.Println("Listening on port 8080....")
	http.ListenAndServe("localhost:8080", nil)
}
