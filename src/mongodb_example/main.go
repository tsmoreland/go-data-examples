package main

import (
	"context"
	"fmt"
	"github.com/tsmoreland/go-data-examples/mongodb_example/configuration"
	"github.com/tsmoreland/go-data-examples/mongodb_example/crud"
	"github.com/tsmoreland/go-data-examples/mongodb_example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func main() {
	config, err := configuration.NewBuilder().
		AddJsonFile("settings.json").
		AddUserSecrets().
		AddEnvironment().
		Build()
	if err != nil {
		log.Fatal(err)
	}

	uri, err := buildConnectionString(config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	collection := client.Database("go-example").Collection("People")

	person := model.Person{FirstName: "John", LastName: "Smith"}
	result, err := crud.Add(collection, context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	id := result.InsertedID
	fmt.Println("Last inserted id ", id)

	filter := bson.D{{"lastName", "Smith"}}
	match, err := crud.FindOneByFilter(collection, context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found: %v", match.FirstName)

	update := bson.D{{"$set",
		bson.D{
			{"firstName", "James"},
		},
	}}

	matches, modified, err := crud.Update(collection, context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Update: %v matched, %v modified", matches, modified)

	deleteCount, err := crud.Delete(collection, context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v records", deleteCount)
}
