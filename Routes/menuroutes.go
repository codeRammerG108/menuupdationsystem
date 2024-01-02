package Routes

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Menu struct {
	MenuID       int
	MenuName     string
	ImageUrl     string
	MenuDesc     string
	MenuPrice    float64
	MenuStock    float64
	MenuDiscount float64
	UserID       int
}

func MenuRoutes(app *fiber.App, DataBase *sql.DB) {
	DB = DataBase
	app.Use(cors.New())
	routeGroup := app.Group("/menu")

	routeGroup.Post("/readMenu", fetchMenuData)

	routeGroup.Post("/addMenu", addMenu)
}

func fetchMenuData(c *fiber.Ctx) error {
	var userId struct {
		ID int `json:"user_id"`
	}
	if err := c.BodyParser(&userId); err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}
	rows, err := DB.Query("SELECT menu_id, name, image_url, description, price, stock, discount, user_id FROM menus where user_id = $1", userId.ID)

	if err != nil {
		log.Println("Error executing query:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	var data []Menu

	// Iterate over the rows and scan data into User struct
	for rows.Next() {
		var menu Menu
		if err := rows.Scan(&menu.MenuID, &menu.MenuName, &menu.ImageUrl, &menu.MenuDesc, &menu.MenuPrice, &menu.MenuStock, &menu.MenuDiscount, &menu.UserID); err != nil {
			log.Println("Error scanning row:", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		data = append(data, menu)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return fetched data as JSON response
	return c.JSON(data)
}

func addMenu(c *fiber.Ctx) error {
	var newmenu struct {
		MenuName     string
		ImageUrl     string
		MenuDesc     string
		MenuPrice    float64
		MenuStock    float64
		MenuDiscount float64
		UserID       int
	}
	if err := c.BodyParser(&newmenu); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := DB.Exec("INSERT INTO menus (name, image_url, description, price, stock, discount, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		newmenu.MenuName, newmenu.ImageUrl, newmenu.MenuDesc, newmenu.MenuPrice, newmenu.MenuStock, newmenu.MenuDiscount, newmenu.UserID)

	if err != nil {
		log.Println("Error executing query:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	} else {
		return c.JSON("Menu Added")
	}
}
