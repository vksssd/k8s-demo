package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var client *mongo.Client

func handler(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("testdb").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	var result struct {
		Name  string  `bson:"name"`
		Value float64 `bson:"value"`
	}
	collection.FindOne(ctx, bson.M{"name": "pi"}).Decode(&result)
	fmt.Fprintf(w, "Hello, World! Pi: %v", result.Value)
}

func main() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}