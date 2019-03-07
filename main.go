package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("testing").Collection("numbers")

	if collection != nil {
		fmt.Print("collection created")
	}

}
