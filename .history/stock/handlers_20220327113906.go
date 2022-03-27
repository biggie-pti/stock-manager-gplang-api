// stock/handlers.go
package stock

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/jinzhu/gorm"
)

type ProductHandler struct {
    repository *ProductRepository
}

func (handler *ProductHandler) GetAll(c *fiber.Ctx) error {
    var products []Product = handler.repository.FindAll()
    return c.JSON(products)
}

func (handler *ProductHandler) Get(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    product, err := handler.repository.Find(id)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "status": 404,
            "error":  err,
        })
    }

    return c.JSON(product)
}

func (handler *ProductHandler) Create(c *fiber.Ctx) error {
    data := new(Product)

    if err := c.BodyParser(data); err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
    }

    item, err := handler.repository.Create(*data)

    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  400,
            "message": "Failed creating item",
            "error":   err,
        })
    }

    return c.JSON(item)
}

func (handler *ProductHandler) Update(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))

    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  400,
            "message": "Item not found",
            "error":   err,
        })
    }

    product, err := handler.repository.Find(id)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "message": "Item not found",
        })
    }

    productData := new(Product)

    if err := c.BodyParser(productData); err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }

    product.Name = productData.Name
    product.Description = productData.Description
    pro.Status = todoData.Status

    item, err := handler.repository.Save(todo)

    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Error updating todo",
            "error":   err,
        })
    }

    return c.JSON(item)
}

func (handler *TodoHandler) Delete(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  400,
            "message": "Failed deleting todo",
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

func NewTodoHandler(repository *TodoRepository) *TodoHandler {
    return &TodoHandler{
        repository: repository,
    }
}

func Register(router fiber.Router, database *gorm.DB) {
    database.AutoMigrate(&Todo{})
    todoRepository := NewTodoRepository(database)
    todoHandler := NewTodoHandler(todoRepository)

    movieRouter := router.Group("/todo")
    movieRouter.Get("/", todoHandler.GetAll)
    movieRouter.Get("/:id", todoHandler.Get)
    movieRouter.Put("/:id", todoHandler.Update)
    movieRouter.Post("/", todoHandler.Create)
    movieRouter.Delete("/:id", todoHandler.Delete)
}
