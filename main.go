package main

import (
	"log"
	"path"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New(fiber.Config{
		Prefork: true,
		GETOnly: true,
	})

	s.Get("/dl/:file", func(c *fiber.Ctx) error {
		file := c.Params("file")
		path := path.Join("dls", file)

		return c.SendFile(path)
	})

	s.Get("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(s.Listen(":3000"))
}
