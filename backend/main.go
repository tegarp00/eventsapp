package main

import (
	"event/handler"
	"event/storage"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type repository struct {
  dbase *gorm.DB
}

func (r *repository) SetupRoutes(app *fiber.App) {
  api := app.Group("/api/v1")

  api.Get("/", handler.Test)
}

func main() {

	// setup config
  err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}


  db, err := storage.NewConnection(config)
  if err != nil {
    log.Fatal("database tidak bisa di load")
  }


  route := repository{
    dbase: db,
  }

  // setup fiber(http route)
  app := fiber.New()
  route.SetupRoutes(app)
  app.Listen(":" + os.Getenv("PORT"))

}
