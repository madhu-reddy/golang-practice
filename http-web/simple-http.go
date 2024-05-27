package main

import "net/http"

func main() {
	//http.ListenAndServe("localhost:3000", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe("localhost:3001", mux)

}
