package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sms2sakthivel/product-manager/products/model"
	"github.com/sms2sakthivel/product-manager/products/service"
)

func RegisterRoutes(app *fiber.App, service *service.ProductService) {
	app.Get("/", ProductServiceInfo)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/products", func(c *fiber.Ctx) error { return GetAllProducts(c, service) })
	app.Get("/products/:id", func(c *fiber.Ctx) error { return GetProductByID(c, service) })
	app.Post("/products", func(c *fiber.Ctx) error { return CreateProduct(c, service) })
	app.Put("/products/:id", func(c *fiber.Ctx) error { return UpdateProduct(c, service) })
	app.Delete("/products/:id", func(c *fiber.Ctx) error { return DeleteProduct(c, service) })
}

// ProductServiceInfo returns information about the Product Service
//
// @Summary      Product Service Info
// @Description  Returns basic information about the Product Service
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func ProductServiceInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Product Service"})
}

// GetAllProducts retrieves all products
//
// @Summary      Get All Products
// @Description  Retrieve a list of all products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Product
// @Failure      500  {object}  fiber.Error
// @Router       /products [get]
func GetAllProducts(c *fiber.Ctx, service *service.ProductService) error {
	products, err := service.GetProducts()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(products)
}

// GetProductByID retrieves a product by their ID
//
// @Summary      Get Product by ID
// @Description  Retrieve a product by their ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  model.Product
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /products/{id} [get]
func GetProductByID(c *fiber.Ctx, service *service.ProductService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := service.GetProduct(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}
	return c.JSON(product)
}

// CreateProduct adds a new product
//
// @Summary      Create a New Product
// @Description  Add a new product to the system
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Product details"
// @Success      201   {object}  model.Product
// @Failure      400   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /products [post]
func CreateProduct(c *fiber.Ctx, service *service.ProductService) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	if err := service.CreateProduct(&product); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

// UpdateProduct modifies details of an existing product
//
// @Summary      Update an Existing Product
// @Description  Modify details of an existing product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Product ID"
// @Param        product  body      model.Product  true  "Updated product details"
// @Success      200   {object}  model.Product
// @Failure      400   {object}  fiber.Error
// @Failure      404   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /products/{id} [put]
func UpdateProduct(c *fiber.Ctx, service *service.ProductService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	product.ID = uint(id)
	if err := service.UpdateProduct(&product); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(product)
}

// DeleteProduct removes a product by their ID
//
// @Summary      Delete a Product
// @Description  Remove a product by their ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Product ID"
// @Success      204
// @Failure      500  {object}  fiber.Error
// @Router       /products/{id} [delete]
func DeleteProduct(c *fiber.Ctx, service *service.ProductService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteProduct(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
