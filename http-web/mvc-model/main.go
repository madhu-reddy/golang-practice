package main

import (
	"mvcmodel/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
