package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World\n"))
	//fmt.Fprintf(w, "Method : %s, Host : %s\n", r.Method, r.Host)
	//fmt.Fprintf(w, "Path : %s, Query : %s\n", r.URL.Path, r.URL.Query())
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "%s", body)
}

/*
A multiplexer
A router that matches the request URI to a handler function

One or more handlers
Functions that handle the requests and return the responses

The multiplexer is quite straightforward. It simply matches the request URI to a handler according to a URL route. For example, you want to match the URL route /home to a homeHandler function.

The handler function is where the real work is done. It takes in a request, does some processing with the data from the request, and returns a response.

In this case, it (multiplexer) is left as nil, which means it will assume the default multiplexer, http.DefaultServeMux. http.DefaultServeMux is an instance of the http.ServeMux struct, which in turn, implements the http.Handler interface.

The http.ServeMux multiplexer has a HandleFunc method that registers a function as a handler for a given URL pattern. This HandleFunc method takes in a URL pattern and a function as parameters. The function must have the signature func(http.ResponseWriter, *http.Request) and will be called when a request matches the URL pattern.
*/
