package DBConnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DBinit() (*sql.DB, error) {
	fmt.Printf("Starting with the connection")

	host := "localhost"
	port := "5433"
	user := "postgres"
	password := "Admin123"
	dbname := "menuupdationsystem"

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return db, nil

}
