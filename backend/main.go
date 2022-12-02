package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// set config untuk membaca variabel
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not reading")
	}

	// mengkoneksikan database pakai gorm
	dsn := viper.GetString("server.DB_URL")
	_, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("db Connected!")


  app := fiber.New()
  
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello World")
  })

  app.Listen(":" + viper.GetString("server.PORT"))
}
