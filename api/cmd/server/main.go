package main

import (
	"net/http"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/registry"
)

func main() {
	h, err := registry.NewRegistry()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", h)
}
