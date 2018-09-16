package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../../html/*")
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"Title": "Hello world"})
	})
	r.POST("/post", posting)
	r.Run(":8080")
}

func posting(c *gin.Context) {
	name := c.PostForm("name")
	message := c.PostForm("message")
	c.HTML(http.StatusOK, "index.html", gin.H{"Title": "postOK"})
	fmt.Printf("name: %s;message: %s\n", name, message)
}
