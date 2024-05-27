package controller

import (
	"encoding/json"
	"mvcmodel/views"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

		}
	}
}
