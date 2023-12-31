package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		var body interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			return
		}

		if body != nil {
			fmt.Println(body)
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
