package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	fmt.Println("Listening to port 9005")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":9005", r))

}

func YourHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Gorilla with mongodb!\n"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoserver:27017"))

	errConnect := client.Connect(ctx)
	if errConnect != nil {
		fmt.Println("ops !! unable to connect ")
	}

	collection := client.Database("testing").Collection("numbers")

	if collection != nil {
		fmt.Println("collection created")
	}

	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	if err != nil {
		fmt.Println("ops !! unable insert record ")
	}

	id := res.InsertedID

	fmt.Println("inserted id ", id)

}
