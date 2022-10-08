package main

import (
	"context"
	"github.com/tsmoreland/go-data-examples/mongodb_example/configuration"
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

}
