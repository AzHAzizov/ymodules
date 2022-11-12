package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	fiber2 "github.com/gofiber/fiber/v2"
)

func main() {


	f1 := fiber.New()
	f2 := fiber2.New()

	fmt.Println(f1)
	fmt.Println(f2)
}
