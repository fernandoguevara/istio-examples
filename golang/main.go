package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/notes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"anywhere": "v2",
		})
	})

	r.GET("/notes/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user": "v2",
		})
	})

	r.GET("/notes/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"admin": "v2",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
