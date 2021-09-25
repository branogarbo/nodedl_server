package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New(fiber.Config{
		Prefork: true,
		GETOnly: true,
	})

	s.Get("/dl/:file", func(c *fiber.Ctx) error {
		file := c.Params("file")

		return c.SendFile("./dls/" + file)
	})

	log.Fatal(s.Listen("172.20.10.4:3000"))
}
