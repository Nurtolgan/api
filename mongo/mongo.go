package mongo

import (
	"context"
	"fmt"

	"api/debugger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	debugger.CheckError("ConnectToMongo", err)

	return client, nil
}

func CreateUserHandler(cv Cv) error {
	client, err := connectToMongo()
	debugger.CheckError("CreateUserHandler", err)
	result, err := client.Database("Vladimir").Collection("Cv").InsertOne(context.Background(), cv)
	debugger.CheckError("Error CreateUserHandler", err)
	fmt.Println(result)
	return nil
}

func GetCVByUsername(username string) (Cv, error) {
	client, err := connectToMongo()
	debugger.CheckError("GetCVByUsername", err)

	var cv Cv
	err = client.Database("Vladimir").Collection("Cv").FindOne(context.Background(), bson.M{"user.username": username}).Decode(&cv)
	debugger.CheckError("Error GetCVByUsername", err)

	return cv, nil
}
