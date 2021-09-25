package main

import (
	"log"
	"net"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	addrs, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	var privateIP string

	for _, a := range addrs {
		if !a.IsLoopback() && a.To4() != nil {
			privateIP = a.String()
			break
		}
	}

	//////////////////////////////////////////////

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

	log.Fatal(s.Listen(privateIP + ":3000"))
}
