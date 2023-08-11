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

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	UserName string  `json:"username"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Address  Address `json:"address"`
	Company  Company `json:"company"`
}

type Users []User

func GetTodos() Todos {
	resp, err := http.Get(jsonplaceholderApi + "todos")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result Todos
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func GetUsers() Users {
	resp, err := http.Get(jsonplaceholderApi + "users")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result Users
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func GetUserByID(id int) User {
	resp, err := http.Get(jsonplaceholderApi + "users/" + strconv.Itoa(id))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result User
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

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		userIDParam := c.Params("id")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid ID")
		}

		user := GetUserByID(userID)

		return c.Render("user", fiber.Map{
			"User": user,
		})
	})

	app.Listen(":3000")
}
