package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func getTodosHandler(c *gin.Context) {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t

	c.JSON(http.StatusCreated, "created todo.")
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.POST("/todos", createTodosHandler)

	r.Run(":1234")
}
