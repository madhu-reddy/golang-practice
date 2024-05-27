1) Entrypoint main.go
2) mux in the main.go, we are getting from the controller layer "route.go" file register() function.
3) The register() function  returns the "mux", when it returns the "mux", it will attach all the endpoints (/ping, /) that we defined.
4) Then we create the endpoints ping.go, crud.go in the "controller" package.


