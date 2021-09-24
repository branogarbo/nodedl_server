package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New(fiber.Config{
		Prefork: true,
		GETOnly: true,
	})

	dl := s.Group("/dl")

	dl.Get("/:file", func(c *fiber.Ctx) error {
		file := c.Params("file")

		return c.SendFile("./dls/" + file)
	})

	s.Listen("172.20.10.4:3000")
}
