package mongo

import (
	"context"
	"fmt"

	"api/debugger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println(cv)
	return cv, nil
}

func DeleteUserById(id string) error {
	client, err := connectToMongo()
	debugger.CheckError("DeleteUserByUsername", err)

	objectid, err := primitive.ObjectIDFromHex(id)
	debugger.CheckError("Error OI from hex", err)

	result, err := client.Database("Vladimir").Collection("Cv").DeleteOne(context.Background(), bson.M{"_id": objectid})
	debugger.CheckError("Error DeleteUserByUsername", err)

	fmt.Println(result)
	return nil
}

func UpdateUserById(id string, cv Cv) error {
	client, err := connectToMongo()
	debugger.CheckError("UpdateUserById", err)

	objectID, err := primitive.ObjectIDFromHex(id)
	debugger.CheckError("Error creating ObjectID", err)

	var filter = bson.M{"_id": objectID}
	var update = bson.M{"$set": cv}

	result, err := client.Database("Vladimir").Collection("Cv").UpdateOne(context.Background(), filter, update)
	debugger.CheckError("Error updating document", err)

	fmt.Printf("Updated %v Document!\n", result.ModifiedCount)

	return nil
}

func GetAllCvsByQuery(args ...string) ([]Cv, error) {
	client, err := connectToMongo()
	debugger.CheckError("GetAllCvsByQuery", err)

	var filter bson.M

	if args[0] != "" || args[1] != "" || args[2] != "" || args[3] != "" {
		filter = bson.M{
			"$and": []bson.M{},
		}

		if args[0] != "" {
			filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"user.username": args[0]})
		}
		if args[1] != "" {
			filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"user.contacts.city": args[1]})
		}
		if args[2] != "" {
			filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"user.baseinfo.birthdaydate": args[2]})
		}
		if args[3] != "" {
			filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"user.special.careerobjective": args[3]})
		}
	} else {
		filter = bson.M{}
	}

	cur, err := client.Database("Vladimir").Collection("Cv").Find(context.Background(), filter)
	debugger.CheckError("Error GetAllCvsByQuery", err)

	var cvs []Cv
	for cur.Next(context.Background()) {
		var cv Cv
		err := cur.Decode(&cv)
		debugger.CheckError("Error GetAllCvsByQuery", err)
		cvs = append(cvs, cv)
	}

	return cvs, nil
}
