package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const jsonplaceholderApi = "https://jsonplaceholder.typicode.com/"

type TodoResponse []struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserResponse []User

func GetTodos() TodoResponse {
	resp, err := http.Get(jsonplaceholderApi + "todos")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result TodoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func GetUsers() UserResponse {
	resp, err := http.Get(jsonplaceholderApi + "users")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result UserResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func GetUserByID(id int, users UserResponse) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		values := GetTodos()

		return c.Render("todos", fiber.Map{
			"Results": values,
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		values := GetUsers()

		return c.Render("users", fiber.Map{
			"Results": values,
		})
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		users := GetUsers()

		userIDParam := c.Params("id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid ID")
		}

		user := GetUserByID(userID, users)
		if user == nil {
			return c.Status(http.StatusNotFound).SendString("User not found")
		}

		return c.Render("user", fiber.Map{
			"User": user,
		})
	})

	app.Listen(":3000")
}
