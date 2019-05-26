package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(
			200,
			"index.html",
			gin.H{
				"msg": "Hello World",
			},
			)

	})

	r.Run(":8088")
}
