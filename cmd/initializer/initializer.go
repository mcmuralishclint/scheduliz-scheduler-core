package initializer

import (
	"context"
	"log"
	"scheduler-service/config"
	"scheduler-service/repo"
)

var mongoStore *repo.MongoStore // Use a store, not just the client
var cfg = config.New()

func DbInit() {
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
}
