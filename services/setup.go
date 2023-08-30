package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx                  context.Context
	tasksCollection      *mongo.Collection
	workspacesCollection *mongo.Collection
)

func init() {
	var err error
	user, password, dbName := os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("DB_NAME")
	ctx = context.TODO()
	link := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.5k1ygzv.mongodb.net/", user, password)
	mongoconn := options.Client().ApplyURI(link)
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo: ", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	tasksCollection = mongoclient.Database(dbName).Collection("tasks")
	workspacesCollection = mongoclient.Database(dbName).Collection("workspaces")
}
