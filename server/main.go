package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os" 
	"github.com/ivanes285/GOReactMongo/models"  // forma de importar un paquete local (models) tome en cuenta la ruta del paquete (models) 
)

func main() {
	app := fiber.New()                // Creamos una instancia de la aplicación de Fiber
	app.Static("/", "../client/dist") // Establecemos los archivos estáticos para el frontend en este caso desde (React)

	//SETTING CORS
	app.Use(cors.New(cors.Config{ // Configuración de CORS para permitir el acceso a la API desde cualquier origen
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// VALIDATION .ENV
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// CONNECT TO MONGODB  (ATLAS)
	MONGOURI := os.Getenv("MONGODB_URI")
	if MONGOURI == "" {
		log.Fatal("Debes establecer una variable de entorno llamada 'MONGODB_URI' para establecer la conección")
	}
	//context es un objeto que se utiliza para cancelar operaciones o para establecer un tiempo de espera
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGOURI))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")

	// defer permite ejecutar una función al finalizar el programa (en este caso la desconexión de la base de datos)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}() // () para ejecutar la función

	coll := client.Database("goreactmongo").Collection("users")

	// INSERTAR UN DOCUMENTO
	coll.InsertOne(context.TODO(), bson.M{"name": "John"})



	//ROUTES
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get users",
		})
	})

	app.Post("/api/v1/users", func(c *fiber.Ctx) error {

		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
	      
		return c.JSON(fiber.Map{
			"data": "Guardando user",
		})
	})


	
	PORT := os.Getenv("PORT") // Obtenemos el puerto de la variable de entorno PORT
	if PORT == "" {
		PORT = "4000"
	}

	//SERVER
	log.Fatal(app.Listen(":" + PORT)) // Iniciamos el servidor y si hay un error lo mostramos en la consola
	fmt.Println("Server is running on port", PORT)
}
