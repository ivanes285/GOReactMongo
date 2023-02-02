package main

import (

	"fmt"
	"log"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ivanes285/GOReactMongo/db"
	"github.com/ivanes285/GOReactMongo/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := run()    // ejecutamos la función run y la asignamos a la variable err
	if err != nil { // si err es diferente de nil entonces mostramos el error en la consola
		panic(err)
	}
}

func run() error {

	// VALIDATION .ENV
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// DATABASE
	err = db.ConnectDB() // Conectamos a la base de datos
	if err != nil {
		return err
	}

	defer db.CloseDB() // Defer permite ejecutar una función al final en este caso para cerrar la conexión a la base de datos

	// SERVER
	app := fiber.New() // Creamos una instancia de la librería Fiber para crear el servidor

	app.Get("*", func(c *fiber.Ctx) error { // Creamos una ruta para que cuando se ingrese a una ruta que no exista se envíe el index.htm
		return c.SendFile("../dist/index.html")
	})

	app.Get("*", func(c *fiber.Ctx) error { // Creamos una ruta para que cuando se ingrese a una ruta que no exista se envíe el index.htm
		return c.SendString("No puede desde aqui")
	})

 
	// STATIC FILES (REACT)
	app.Static("/", "../dist") // Establecemos los archivos estáticos para el frontend en este caso desde (React)
		// STATIC FILES (REACT)
	app.Static("*", "../dist") // Establecemos los archivos estáticos para el frontend en este caso desde (React)
	
	
	

	// MIDDLEWARES
	app.Use(logger.New())         // logger permite mostrar en la consola las peticiones que se hacen a la API
	app.Use(recover.New())        // recover permite mostrar en la consola los errores y no se caiga el servidor en el caso que se ejecute un panic
	app.Use(cors.New(cors.Config{ // Configuración de CORS para permitir el acceso a la API desde cualquier origen
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept,Access-Control-Allow-Origin",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// ROUTES
	routes.AddUsersGroup(app) // Agregamos las rutas de los usuarios


	// PORT
	PORT := os.Getenv("PORT") // Obtenemos el puerto de la variable de entorno PORT
	if PORT == "" {
		PORT = "4000"
	}

	//SERVER
	log.Fatal(app.Listen(":" + PORT)) // Iniciamos el servidor y si hay un error lo mostramos en la consola
	fmt.Println("Server is running on port", PORT)
	return nil // retornamos nil porque no hay error
}
