package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	server := ""
	intent := ""
	port := ""

	fmt.Println(len(os.Args))

	if len(os.Args) >= 3 {
		server = os.Args[1]
		port = os.Args[2]
		intent = os.Args[3]
	}

	if server == "" {

		r := mux.NewRouter()
		// Routes consist of a path and a handler function.
		r.HandleFunc("/", YourHandler)

		fmt.Println("Listening to port 9005")
		// Bind to a port and pass our router in
		log.Fatal(http.ListenAndServe(":9005", r))
	} else {

		if intent == "c" {
			fmt.Println("Creating connection to mongo db")
			createConnection(server, port)
		} else {
			createRecord(server, port)
		}
	}

}

func YourHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Gorilla with mongodb!\n"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoserver:27017"))

	errConnect := client.Connect(ctx)

	errPing := client.Ping(ctx, readpref.Primary())

	if errPing != nil {
		fmt.Println("opss. ... unable to ping")
	} else {
		fmt.Println("successful connection!")
	}

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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla with mongodb!\n"))
}

func createRecord(server string, port string) {

	targetConnection := "mongodb://" + server + ":" + port
	fmt.Println("using the following connection string:", targetConnection)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.NewClient(options.Client().ApplyURI(targetConnection))

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

func createConnection(server string, port string) {

	targetConnection := "mongodb://" + server + ":" + port
	fmt.Println("using the following connection string:", targetConnection)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.NewClient(options.Client().ApplyURI(targetConnection))

	errConnect := client.Connect(ctx)

	errPing := client.Ping(ctx, readpref.Primary())

	if errPing != nil {
		fmt.Println("opss. ... unable to ping")
	} else {
		fmt.Println("successful connection!")
	}

	if errConnect != nil {
		fmt.Println("ops !! unable to connect ")
	}
}
