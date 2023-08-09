package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const jsonplaceholderApi = "https://jsonplaceholder.typicode.com/"

type TodoResponse []struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type UserResponse []struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

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

	app.Listen(":3000")
}
