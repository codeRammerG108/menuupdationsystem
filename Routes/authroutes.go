package Routes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var DB *sql.DB

type User struct {
	ID       int
	Username string
	Password string
}

// GetAuthRoutes is an exported function that sets up and returns Fiber app with authentication routes
func GetAuthRoutes(app *fiber.App, DataBase *sql.DB) {
	DB = DataBase
	app.Use(cors.New())
	routeGroup := app.Group("")

	routeGroup.Get("/readAllUsers", fetchDataHandler)

	routeGroup.Get("/authUser", authUser)

	routeGroup.Post("/addNewUser", addNewUser)

}

func fetchDataHandler(c *fiber.Ctx) error {
	// Fetch data from the "users" table in the database
	rows, err := DB.Query("SELECT id, username, password FROM users")
	if err != nil {
		log.Println("Error executing query:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	var data []User

	// Iterate over the rows and scan data into User struct
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			log.Println("Error scanning row:", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		data = append(data, user)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return fetched data as JSON response
	return c.JSON(data)
}

func authUser(c *fiber.Ctx) error {

	// Retrieve parameters from the query string
	username := c.Query("username")
	password := c.Query("password")

	// Your logic here

	rows, err := DB.Query("SELECT id, username, password FROM users")
	if err != nil {
		log.Println("Error executing query:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		var Id int
		var Username, Password string

		err := rows.Scan(&Id, &Username, &Password)
		if err != nil {
			fmt.Println("Error in Scanning rows: ", err)
		}

		if username == Username && password == Password {
			return c.JSON(Id)
		}
	}
	return c.JSON(0)
}

func addNewUser(c *fiber.Ctx) error {
	var newUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&newUser); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", newUser.Username, newUser.Password)
	if err != nil {
		log.Println("Error executing query:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	} else {
		return c.JSON("Added new user")
	}
}
