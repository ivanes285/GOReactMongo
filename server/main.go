package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()                // Creamos una instancia de la aplicación de Fiber
	app.Static("/", "../client/dist") // Establecemos la ruta de la carpeta de los archivos estáticos para el frontend en este caso desde (React)

	app.Use(cors.New(cors.Config{   // Configuración de CORS para permitir el acceso a la API desde cualquier origen
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//Routes
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Get("/api/v1/tasks", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "tasks",
		})
	})

	app.Get("/api/v1/ss", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "tasks",
		})
	})

	PORT:= os.Getenv("PORT")                      // Obtenemos el puerto de la variable de entorno PORT
	if PORT == "" {
	  PORT = "3000"
	}
	log.Fatal(app.Listen(":" + PORT))      // Iniciamos el servidor y si hay un error lo mostramos en la consola
	fmt.Println("Server is running on port", PORT)
}
