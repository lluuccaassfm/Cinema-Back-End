package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"lucas-lima/cinema/configs"
	"time"
)

var Client *mongo.Client

var Db *mongo.Database

func CreateClient() {

	Client, err := mongo.NewClient(options.Client().ApplyURI(configs.MONGO_HOST))

	if err != nil {
		log.Println("[FATAL] could not create client for database")
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

	err = Client.Connect(ctx)

	defer cancel()

	if err != nil {
		log.Println("[FATAL] could not connect to database")
		panic(err)
	}

	Db = Client.Database("cine")

	return
}