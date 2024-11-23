package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers for v1 and v2 routes
func loginEndpoint(c *gin.Context) {
	// Example: Parse JSON payload and respond
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Placeholder login logic
	if requestBody.Username == "admin" && requestBody.Password == "password" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": requestBody.Username})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
}

func submitEndpoint(c *gin.Context) {
	// Example: Handle form submission
	var formData struct {
		Data string `json:"data" binding:"required"`
	}

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Placeholder for processing submitted data
	c.JSON(http.StatusOK, gin.H{"message": "Data received", "data": formData.Data})
}

func readEndpoint(c *gin.Context) {
	// Example: Respond with a sample resource
	resource := gin.H{
		"id":      1,
		"content": "This is a sample resource.",
	}
	c.JSON(http.StatusOK, resource)
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

/*
curl -X POST -H "Content-Type: application/json" \
-d '{"title": "New Submission", "description": "This is a test submission."}' \
http://localhost:8080/v1/submit

curl -X POST http://localhost:8080/v1/read
curl -X POST http://localhost:8080/v2/read

curl -X POST -H "Content-Type: application/json" -d '{"username":"admin", "password":"password"}' http://localhost:8080/v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username":"admin", "password":"password"}' http://localhost:8080/v2/login
*/
