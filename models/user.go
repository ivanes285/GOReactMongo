package models

import "go.mongodb.org/mongo-driver/bson/primitive" //para usar el tipo ObjectID
type User struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` //bson es para mongo
	Name string `json:"name"`
}
