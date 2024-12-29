package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	http.ListenAndServe(":8080", router)
}

/*

1) go doc http ListenAndServe
func ListenAndServe(addr string, handler Handler) error
    ListenAndServe listens on the TCP network address addr and then calls
    Serve with handler to handle requests on incoming connections. Accepted
    connections are configured to enable TCP keep-alives.

    The handler is typically nil, in which case DefaultServeMux is used.

    ListenAndServe always returns a non-nil error.


2) go doc http Handler
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
A Handler responds to an HTTP request.


3) go doc gin default
func Default(opts ...OptionFunc) *Engine
    Default returns an Engine instance with the Logger and Recovery middleware
    already attached.


4)
go doc gin engine
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)


5) go doc gin handlerfunc
type HandlerFunc func(*Context)
    HandlerFunc defines the handler used by gin middleware as return value.

6)
router.GET: Registers a route for the GET HTTP method.
Handler Function: A function of type gin.HandlerFunc that takes *gin.Context as an argument.
c.JSON: Sends a JSON response back to the client.
*/
