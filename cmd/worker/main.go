package main

import (
	"fmt"
	"scheduler-service/config"
)

var cfg *config.Config

func main() {
	cfg = config.New()
	//initalizedConfigs := initializer.DbInit(cfg)
	fmt.Println("Started Worker Service")
}
