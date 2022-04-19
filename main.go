package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("mongogo").Collection("Test")
	docs := []interface{}{
		bson.D{{"message","hello"}},
		bson.D{{"message","bye bye"}},
	}
	res, insertErr := collection.InsertMany(ctx, docs);
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println("Number of inserted rows: %d\n", len(res.InsertedIDs))
	cur, currErr := collection.Find(context.TODO(), bson.D{})
	if currErr != nil {
		panic(currErr)
	}
	for cur.Next(context.TODO()) {
		var result bson.D
		if err := cur.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	defer cur.Close(ctx)
}