package main

import (
	"context"
	"log"
	"time"

	"github.com/ij4l/foodCatalog/external/database"
	"github.com/ij4l/foodCatalog/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("Cannot load configuration :", err)
	}

	db, err := database.ConnectMongoDb(config, ctx)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB :", err)
	}

	coll := db.Collection("favorite_books")
	data, err := coll.InsertOne(ctx, User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "testing@gmail.com",
		Password: "password",
	})
	if err != nil {
		log.Fatal("Cannot insert data :", err)
	}

	log.Println("Inserted data:", data.InsertedID)
}