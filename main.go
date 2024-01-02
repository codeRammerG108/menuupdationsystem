package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"

	DBconnection "github.com/codeRammer07/server/DBConnection"
	AllRoute "github.com/codeRammer07/server/Routes"
)

var DB *sql.DB

func main() {
	var err error
	fmt.Println("Server is starting from on port!")

	DB, err := DBconnection.DBinit()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("\nConnection Established")
	}
	defer DB.Close()

	app := fiber.New()

	AllRoute.GetAuthRoutes(app, DB)
	AllRoute.MenuRoutes(app, DB)

	// Start the Fiber app
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
