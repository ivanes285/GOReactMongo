package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func ConnectDB() error {
	MONGOURI := os.Getenv("MONGODB_URI")
	if MONGOURI == "" {
		return errors.New("debes establecer una variable de entorno llamada 'MONGODB_URI' para establecer la conección a la base de datos")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGOURI))

	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB")
	db = client.Database("goreactmongo") //db es una variable global de este paquete (db) y se le asigna el valor de la base de datos goreactmongo que es la que se va a utilizar en el proyecto 

	return nil  // si no hay error retornamos nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
	
}