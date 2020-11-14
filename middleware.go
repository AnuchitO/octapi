package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)
func nong(c *gin.Context) {
	log.Println("nong Handler")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
func helloHandler(c *gin.Context) {
	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func authen(c *gin.Context) {
	log.Println("start middleware")
	c.Next()
	log.Println("end middleware")
}

func main() {
	r := gin.Default()

	r.GET("/login", helloHandler)
	r.Use(authen)
	r.GET("/statements", nong)
	r.Run(":1234")
}
