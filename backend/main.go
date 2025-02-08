package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Define the root route
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hi",
        })
    })

    // Get all users
    r.GET("/users", func(c *gin.Context) {
        users := []string{"Alice", "Bob", "Charlie"}
        c.JSON(http.StatusOK, gin.H{
            "users": users,
        })
    })
  
    // Create a user
    r.POST("/users", func(c *gin.Context) {
        var user map[string]interface{}
        if err := c.ShouldBindJSON(&user); err == nil {
            c.JSON(http.StatusCreated, gin.H{
                "status": "user created",
                "user":   user,
            })
        } else {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
        }
    })

    // Start the server
    r.Run(":8080")
}

