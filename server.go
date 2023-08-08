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

type Response []struct {
	Title string `json:"title"`
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetTodos() Response {
	resp, err := http.Get(jsonplaceholderApi + "todos")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(PrettyPrint(result))

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

	app.Listen(":3000")
}
