package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = EnvMongoURI()
var clientDb *mongo.Client

const dbName = "form-handler-go"

func InitDb() {
	var err error
	clientDb, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = clientDb.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to MongoDB!")
}

func GetMongoDbClient() *mongo.Database {
	return clientDb.Database(dbName)
}
