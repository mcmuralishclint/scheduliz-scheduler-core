package initializer

import (
	"context"
	"log"
	"scheduler-service/config"
	"scheduler-service/repo"
)

var mongoStore *repo.MongoStore

type initializedConfigs struct {
	MongoStore *repo.MongoStore
}

func DbInit(cfg *config.Config) *initializedConfigs {
	var err error
	mongoStore, err = repo.NewClient(cfg.MongoURI) // Assign to global variable
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = mongoStore.Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to Mongo!")
	return &initializedConfigs{MongoStore: mongoStore}
}
