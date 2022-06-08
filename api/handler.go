package api

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Get(c *fiber.Ctx)
	Post(ctx *fiber.Ctx)
	Put(ctx *fiber.Ctx)
	Delete(ctx *fiber.Ctx)
	GetAll(ctx *fiber.Ctx)
}
