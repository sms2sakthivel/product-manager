package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sms2sakthivel/product-manager/products/database"
	"github.com/sms2sakthivel/product-manager/products/model"
	"github.com/sms2sakthivel/product-manager/products/repository"
	"github.com/sms2sakthivel/product-manager/products/routes"
	"github.com/sms2sakthivel/product-manager/products/service"
)

func NewApp() *fiber.App {
	// Step 1: Connect to the database
	database.Connect()

	// Step 2: Auto-migrate Product schema
	database.DB.AutoMigrate(&model.Product{})

	// Step 3: Initialize repository, service, and app
	repo := &repository.GormProductRepository{DB: database.DB}
	service := &service.ProductService{Repo: repo}
	app := fiber.New()

	// Step 4: Register routes
	routes.RegisterRoutes(app, service)
	return app
}
