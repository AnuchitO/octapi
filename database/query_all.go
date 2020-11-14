package main


import (
	"database/sql"
	"os"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID int
	Title string
	Status string
}
func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()


	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}

	var todos []Todo
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		todos = append(todos, Todo{id, title, status})
	}

	fmt.Printf("query all todos success %#v\n", todos)
}

