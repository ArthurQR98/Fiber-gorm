package main

import (
	"github.com/ArthurQR98/fiber-gorm/api"
	"github.com/ArthurQR98/fiber-gorm/config"
	"github.com/ArthurQR98/fiber-gorm/domain"
	"github.com/ArthurQR98/fiber-gorm/repository"
	"github.com/ArthurQR98/fiber-gorm/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	db, _ := domain.InitDB(conf.Database.DNS)
	userRepo, _ := repository.NewPostgresRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewHandler(userService)

	r := fiber.New()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	r.Get("/users/:id", userHandler.Get)
	r.Post("/users", userHandler.Post)
	r.Delete("/users/:id", userHandler.Delete)
	r.Get("/users", userHandler.GetAll)
	r.Put("/users/:id", userHandler.Put)
	r.Listen(":" + conf.Server.Port)
}
