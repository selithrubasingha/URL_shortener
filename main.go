package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "Hello Go URL Shortener !",
		})
})

	err := r.Run(":9808")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}


}