package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()

	// Basic server code
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hello go"})
	// })

	// Form handling code
	r.LoadHTMLGlob("views/*")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)
	r.Run()
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(200, gin.H{"Color: ": fakeForm.Colors})
}
