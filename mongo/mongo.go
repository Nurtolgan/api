package mongo

import (
	"context"
	"fmt"

	"api/debugger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	debugger.CheckError("ConnectToMongo", err)

	return client, nil
}

// func CreateUserHandler(client *mongo.Client, cv *Cv) {
// 	result, err := client.Database("Vladimir").Collection("Cvs").InsertOne(context.Background(), cv)
// 	debugger.CheckError("Insert One", err)
// 	fmt.Println(result)
// }


func CreateUserHandler(cv Cv) {

	client, err := connectToMongo()
	debugger.CheckError("CreateUserHandler", err)
	result, err := client.Database("Vladimir").Collection("Cv").InsertOne(context.Background(), cv)
	debugger.CheckError("Insert One", err)
	fmt.Println(result)
}


