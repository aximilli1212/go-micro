package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "50001"
	mongoURL = "mongodb://mongo:27017"
)

func connectToMongo() (*mongo.Client, error) {
	//connect to the database
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	//connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to mongo: ", err)
		return nil, err
	}

	return c, nil
}

func main() {
	//connect to the database
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client := mongoClient

	fmt.Printf("Connected to mongo: %v\n", client)

}
