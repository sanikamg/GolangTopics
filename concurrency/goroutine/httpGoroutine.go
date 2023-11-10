package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/first", First)
	router.GET("/second", Second)

	router.Run(":8080")

}

func First(c *gin.Context) {
	msg := "hello"
	c.JSON(200, msg)
}

func Second(c *gin.Context) {
	msg := "hiiii"
	c.JSON(200, msg)
}
