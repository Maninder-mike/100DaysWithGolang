package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil{
        log.Fatal(err)
    }
    collection := client.Database("testing").Collection("numbers")
    fmt.Printf("%v, %T", collection, collection)
}
