package config

import "os"

type Config struct {
	ServerPort  string
	MongoURI    string
	RabbitMQURI string
}

func New() *Config {
	return &Config{
		ServerPort:  os.Getenv("SERVER_PORT"),
		MongoURI:    os.Getenv("MONGO_URI"),
		RabbitMQURI: os.Getenv("RABBITMQ_URI"),
	}
}
