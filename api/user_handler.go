package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neogan74y/hot_res/types"
)

func HadlerGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Arina",
		LastName:  "Gushchina",
	}
	return c.JSON(u)
}

func HadlerGetUserById(c *fiber.Ctx) error {
	return c.JSON("123")
}
