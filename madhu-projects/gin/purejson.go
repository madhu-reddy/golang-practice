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
2024-11-23-10-26-00.png
*/
