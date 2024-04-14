package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	mux.Get("/", index)
	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

/*

 It is quite common to use more sophisticated and performant multiplexers offered by third-party packages; for example, the chi package.
 As you can see, the main change is that you need to create a new multiplexer, mux, and this will be used instead of the default multiplexer when you start the server with http.ListenAndServe. Third-party multiplexers are more optimal and capable. In most cases, you should use one unless you are writing something that only uses the standard library.

*/
