package api

import (
	"github.com/ArthurQR98/fiber-gorm/domain"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userService domain.Service
}

func (h *handler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ocurrio un error"})
	}
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ocurrio un error"})
	}
	p, err := h.userService.Find(id)
	if err != nil {
		return ctx.Status(404).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Post(ctx *fiber.Ctx) error {
	u := &domain.User{}
	if err := ctx.BodyParser(&u); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	err := h.userService.Store(u)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&u)
}

func (h *handler) Put(ctx *fiber.Ctx) error {
	u := &domain.User{}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ocurrio un error"})
	}
	if err = ctx.BodyParser(&u); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	err = h.userService.Update(id, u)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&u)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ocurrio un error"})
	}
	err = h.userService.Delete(id)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.Status(201).JSON(fiber.Map{"message": "ok"})
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	p, err := h.userService.FindAll()
	if err != nil {
		return ctx.Status(404).JSON(nil)
	}
	return ctx.JSON(&p)
}

// NewHandler  New handler instantiates a http handler for our user service
func NewHandler(userService domain.Service) *handler {
	return &handler{userService: userService}
}
