// main.go
package main

import (
	"log"

	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/biggie-pti/stock-manager-gplang-api/database"
	"github.com/biggie-pti/stock-manager-gplang-api/stock"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")
	stock.Register(api, database.DB)

	log.Fatal(app.Listen(":5000"))
}
