package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"scheduler-service/model"
)

type MongoStore struct {
	Client *mongo.Client
}

func NewClient(uri string) (*MongoStore, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoStore{Client: client}, nil
}

func (store *MongoStore) ListSchedules() ([]model.Schedule, error) {
	schedules := make([]model.Schedule, 0)
	collection := store.Client.Database("schedule-manager").Collection("schedules")

	// Finding multiple documents returns a cursor
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor
	for cursor.Next(context.TODO()) {
		var schedule model.Schedule
		err := cursor.Decode(&schedule)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (store *MongoStore) AddSchedule(schedule model.Schedule) error {
	collection := store.Client.Database("schedule-manager").Collection("schedules")
	schedule.State = model.Init
	_, err := collection.InsertOne(context.TODO(), schedule)
	return err
}

func (store *MongoStore) GetSchedule(id string) (model.Schedule, error) {
	var schedule model.Schedule
	collection := store.Client.Database("schedule-manager").Collection("schedules")

	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&schedule)
	if err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (store *MongoStore) UpdateSchedule(schedule model.Schedule) error {
	collection := store.Client.Database("schedule-manager").Collection("schedules")
	filter := bson.M{"_id": schedule.ID}
	update := bson.M{"$set": bson.M{
		"name":    schedule.Action,
		"api_key": schedule.Schedule,
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (store *MongoStore) DeleteSchedule(id string) error {
	collection := store.Client.Database("schedule-manager").Collection("schedules")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
