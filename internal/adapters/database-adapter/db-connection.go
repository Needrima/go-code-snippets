package databaseadapter

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB() *mongo.Collection {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONN_STRING"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("connecting to database:", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("pinging database:", err)
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	return collection
}
