package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	router.GET("/purejson", func(ctx *gin.Context) {
		ctx.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	router.Run(":8080")
}

/*
PureJSON serializes the given struct as JSON into the response body. PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
images/purejson/purejson.png

JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
images/purejson/json.png

*/
