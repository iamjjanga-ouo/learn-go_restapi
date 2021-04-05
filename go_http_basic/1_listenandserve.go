// ListenAndServe(addr string, handler Handler) error
//
package main

import (
	"fmt"
	"net/http"
)

type MyWebserverType bool

// Handler interface는 ServeHTTP 메서드를 가지고있다.
// 따로 custom func와 메서드를 구현해서 Handler를 사용할 수 있다.
func (m MyWebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "well, hello there !!!")
	fmt.Fprintf(w, "Request is \n%+v", r)
}

func main() {
	var k MyWebserverType

	http.ListenAndServe("localhost:8080", k)
}
