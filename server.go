package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"os"
)

type Todo struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
}
package database

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("xxxxxx", err)
	}
}
//
//type DB struct {
//	db *sql.DB
//}
//
//func (d *DB) conn() *sql.DB {
//	return db
//}
//
//type Database interface {
//	conn() *DB
//}
//
func Conn() *sql.DB {
	return db
}

func getTodoHandler(c *gin.Context) {
	db := Conn()
	todos, err := queryAll(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusOK, todos)
		return
	}

	items := []Todo{}
	for _, item := range todos {
		if item.Status == status {
			items = append(items, item)
		}
	}

	c.JSON(http.StatusOK, items)
}

type Database interface {
	Prepare(query string) (*sql.Stmt, error)
}

func queryAll(db Database) ([]Todo, error) {
	stmt, err := db.Prepare("SELECT id,title, status FROM todos")
	if err != nil {
		return []Todo{}, fmt.Errorf("can't prepare query all todos statement: %s", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return []Todo{}, fmt.Errorf("can't query all todos: %s", err)
	}

	var todos []Todo
	for rows.Next() {
		var id int
		var title, status string

		err := rows.Scan(&id, &title, &status)
		if err != nil {
			return []Todo{}, fmt.Errorf("can't Scan row into variable: %s", err)
		}
		todo := Todo{id, title, status}
		todos = append(todos, todo)
	}
	return todos, nil
}

func main() {
	r := gin.Default()
	r.POST("/todos", createTodoHandler)
	r.Run(":1234")
}
