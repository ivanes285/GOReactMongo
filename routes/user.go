package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ivanes285/GOReactMongo/db"     // forma de importar un paquete local (db)
	"github.com/ivanes285/GOReactMongo/models" // forma de importar un paquete local (models) tome en cuenta la ruta del paquete (models)
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUsersGroup(app *fiber.App) {
	userGroup := app.Group("/api/v1/users")
	userGroup.Get("/", getUsers)
	userGroup.Post("/", createUser)
	userGroup.Get("/:id", getUserById)
	userGroup.Put("/:id", updateUser)
	userGroup.Delete("/:id", deleteUser)
}

// GET ALL USERS
func getUsers(c *fiber.Ctx) error {

	coll := db.GetDBCollection("users") // obtenemos la coleccion de la base de datos goreactmongo que se encuentra en el paquete db

	var users []models.User                          //creamos un slice(array de tamño dinamico) de tipo User
	results, err := coll.Find(c.Context(), bson.M{}) //retorna primero en formato bson
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error()},
		)
	}

	if err = results.All(c.Context(), &users); // aqui recien estamos asignando el resultado a la variable users que es un slice de tipo User
	err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Could not find users"},
		)
	}

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

// CREATE USER
func createUser(c *fiber.Ctx) error {

	user := new(models.User)
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

	result, err := coll.InsertOne(c.Context(), bson.M{"name": user.Name})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Could not insert user,Failed to create user"},
		)
	}
	return c.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

// GET USER BY ID
func getUserById(c *fiber.Ctx) error {
	coll := db.GetDBCollection("users")

	// find the user by id
	id := c.Params("id") //obtenemos el id del parametro
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id) //convertimos el id a un objeto de tipo ObjectID
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	user := models.User{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not find user",
		})
	}

	return c.Status(200).JSON(fiber.Map{"user": user})
}

func updateUser(c *fiber.Ctx) error {
	// validate the body
	b := new(models.User)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	// update the user
	coll := db.GetDBCollection("users")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update book",
			"message": err.Error(),
		})
	}
	message :="User update successfully"
	status := 200

   if result.ModifiedCount == 0 {
	   message = "User not found or already updated"
	   status = 404
   } 
	// return the user
	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}

func deleteUser(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	// delete the user
	coll := db.GetDBCollection("users")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Could not delete user.Failed to delete user",
			"message": err.Error(),
		})
	}

	 message :="User deleted successfully"
     status := 200

	if result.DeletedCount == 0 {
		message = "User not found or already deleted"
		status = 404
	} 

	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
