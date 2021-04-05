//HandlerFunc(f)
// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
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
	http.ListenAndServe("localhost:8080", http.HandlerFunc(myFunc))
}
