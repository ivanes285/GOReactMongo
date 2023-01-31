package routes

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/ivanes285/GOReactMongo/models" // forma de importar un paquete local (models) tome en cuenta la ruta del paquete (models)
	"github.com/ivanes285/GOReactMongo/db"    // forma de importar un paquete local (db) 
	"go.mongodb.org/mongo-driver/bson"

)


func AddUsersGroup(app *fiber.App) {
	userGroup := app.Group("/api/v1/users")
	userGroup.Get("/", getUsers)
	userGroup.Post("/", createUser)
// 	userGroup.Get("/:id", getBook)
// 	userGroup.Put("/:id", updateBook)
// 	userGroup.Delete("/:id", deleteBook)
}



//GET ALL USERS
 func getUsers (c *fiber.Ctx) error {

	coll := db.GetDBCollection("users")  // obtenemos la coleccion de la base de datos goreactmongo que se encuentra en el paquete db

	var users []models.User                             //creamos un slice(array de tamño dinamico) de tipo User
	results, err := coll.Find(context.TODO(), bson.M{}) //retorna primero en formato bson
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error()},
		)
	}

	if err = results.All(context.TODO(), &users); // aqui recien estamos asignando el resultado a la variable users que es un slice de tipo User
	err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Could not find users"},
		)
	}

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}


//CREATE USER
 func createUser (c *fiber.Ctx) error {

	user:= new(models.User)
	if err := c.BodyParser(&user); //? &user es un puntero a la variable user en este caso se le pasa el valor del body y no hay necesidad de usar user = c.BodyParser() y mas bien verificamos si hay un error en el body y lo validamos con el if

	err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "An error has occurred with data,Invalid body"},
		)
	}
	// VALIDATE USER
	if user.Name == "" || user.Name == " " {
		return c.Status(400).JSON(fiber.Map{
			"message": "User is required"},
		)
	}

	coll := db.GetDBCollection("users") // obtenemos la coleccion de la base de datos goreactmongo que se encuentra en el paquete db
	
	result, err := coll.InsertOne(context.TODO(), bson.M{"name": user.Name})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":  err.Error(),
			"message": "Could not insert user,Failed to create user"},
		)
	}
	return c.Status(200).JSON(fiber.Map{
		"data": result,
	})
}