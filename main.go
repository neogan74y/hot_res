package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)
	"github.com/neogan74y/hot_res/api"

func main() {
	listenAddr := flag.String("listenAddr", ":5001", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HadlerGetUsers)
	apiV1.Get("/user/:id", api.HadlerGetUserById)

	log.Fatal(app.Listen(*listenAddr))
}
