package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetUserCollection from MongoDb
func GetUserCollection() *mongo.Collection {

	connectString := os.Getenv("DB_CONNECTION_STRING")
	cluster := os.Getenv("CLUSTER")
	dbName := os.Getenv("DB_NAME")
	dbCollection := os.Getenv("DB_COLLECTION")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	url := "mongodb+srv://" + username + ":" + password + "@" + connectString + "/" + cluster + "?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOption := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbName).Collection(dbCollection)
}
