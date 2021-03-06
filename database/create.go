package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main(){
	db, err := sql.Open("postgres", "postgres://jjqqdegw:bpOkszJWGxy9ihI5iFkkGXMi1UxOSheW@john.db.elephantsql.com:5432/jjqqdegw ")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        title TEXT,
        status TEXT
    );
    `
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("okay")
}
