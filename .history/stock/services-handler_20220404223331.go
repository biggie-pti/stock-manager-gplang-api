// stock/handlers.go
package stock

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type ServicesHandler struct {
	serviceRepository *ServiceRepository
}

func (handler *ServicesHandler) GetAll(c *fiber.Ctx) error {
	var services []Service = handler.serviceRepository.FindAllServices()
	return c.JSON(services)
}

func (handler *ServicesHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	service, err := handler.serviceRepository.FindService(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": 404,
			"error":  err,
		})
	}

	return c.JSON(service)
}

func (handler *ServicesHandler) Create(c *fiber.Ctx) error {
	data := new(Service)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
	}

	item, err := handler.serviceRepository.AddService(*data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed adding a service",
			"error":   err,
		})
	}

	return c.JSON(item)
}

func (handler *ServicesHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Service not found",
			"error":   err,
		})
	}

	service, err := handler.serviceRepository.FindService(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Service not found",
		})
	}

	serviceData := new(Service)

	if err := c.BodyParser(serviceData); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	service.Name = serviceData.Name
	service.Description = productData.Description
	service.Category = productData.Status

	item, err := handler.repository.Save(product)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error updating product",
			"error":   err,
		})
	}

	return c.JSON(item)
}

func (handler *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed deleting product",
			"err":     err,
		})
	}
	RowsAffected := handler.repository.Delete(id)
	statusCode := 204
	if RowsAffected == 0 {
		statusCode = 400
	}
	return c.Status(statusCode).JSON(nil)
}

func NewProductHandler(repository *ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: repository,
	}
}

func Register(router fiber.Router, database *gorm.DB) {
	database.AutoMigrate(&Product{})
	productRepository := NewProductRepository(database)
	productHandler := NewProductHandler(productRepository)

	stockRouter := router.Group("/product")

	stockRouter.Get("/", productHandler.GetAll)
	stockRouter.Get("/:id", productHandler.Get)
	stockRouter.Put("/:id", productHandler.Update)
	stockRouter.Post("/", productHandler.Create)
	stockRouter.Delete("/:id", productHandler.Delete)

}
