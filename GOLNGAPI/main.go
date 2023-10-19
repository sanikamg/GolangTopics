package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

func main() {
	router := gin.Default()
	router.GET("/getusers", getUsers)
	router.Run(":8181")

}

func getUsers(c *gin.Context) {
	c.JSON(200, users)
}
