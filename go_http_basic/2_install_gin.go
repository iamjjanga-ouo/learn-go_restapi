// Gin?
// gin is a simple command line utility for live-reloading Go web applications.
// install : `go get github.com/codegangsta/gin`
// usage : gin --all -i run main.go

package main

import (
	"fmt"
	"net/http"
)

type MyWebserverType bool

func (m MyWebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
		<head>
			Hi
		</head>
		<body>
			<h1>My name is Iamjjanga!</h1>
			<p> gin is installed now</p>
		</body>
	</html>
	`)
}

func main() {
	var k MyWebserverType

	http.ListenAndServe("localhost:8080", k)
}
