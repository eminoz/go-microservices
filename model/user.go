package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Login struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email",bson:"email"`
	Password string             `json:"password"`
}
type Order struct {
	ProductName string `json:"productName"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
}
type UserDal struct {
	Id      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Email   string `json:"email" bson:"email"`
}
type LoginDal struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
	Token string             `json:"token" bson:"token"`
}
