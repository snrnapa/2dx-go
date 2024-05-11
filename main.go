package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "completed user login",
		})
	})

	router.GET("info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get basic ingo",
		})
	})

	router.POST("info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "crate new info",
		})
	})

	router.DELETE("info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete music data info",
		})
	})

	router.PATCH("info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update part of  music data info",
		})
	})

	router.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello 2DX")
		fmt.Println("hello 2DX")

	})

	router.Run(":8080")
}
