package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// creates a router that comes pre packages with a logger (loggin errors and stuff to the terminal)
	r := gin.Default()

	// like the controller . when user visites "/" route , the function inside is called . 
	r.GET("/",func(c *gin.Context) { // this c containes ALL the info of the HTTP request . very importatnt

		// sends 200 request . gin.H nicely formats into JSON response
		c.JSON(200,gin.H{
			"message": "Hello Go URL Shortener !",
		})
})

	// fires up built in go web server on port 9808.
	err := r.Run(":9808")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}


}