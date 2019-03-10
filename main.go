package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	errConnect := client.Connect(ctx)
	if errConnect != nil {
		fmt.Println("ops !! unable to connect ")
	}

	collection := client.Database("testing").Collection("numbers")

	if collection != nil {
		fmt.Print("collection created")
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	if err != nil {
		fmt.Println("ops !! unable insert record ")
	}

	id := res.InsertedID

	fmt.Println("inserted id ", id)

}
