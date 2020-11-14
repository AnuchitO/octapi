package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"time"
)

type Todo struct {
	ID       int       `json:"id"`
	Birthday time.Time `json:"birthday"`
	CreateAt time.Time `json:"createAt"`
}

var todo Todo

func createTodoHandler(c *gin.Context) {
	var t Todo
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("%#v\n", t)
	log.Println(t.Birthday)

	t.CreateAt = time.Now()

	c.JSON(http.StatusCreated, t)
}

func main() {
	r := gin.Default()
	r.POST("/todos", createTodoHandler)
	r.Run(":1234")
}
