package main

import (
	"encoding/json"
	"net/http"
	structs1 "structs/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := structs1.Response{
				Code: http.StatusOK,
				Body: "pong",
			}

			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}

/*
Output
http://localhost:3000/ping
{"Code":200,"Body":"pong"}
*/
