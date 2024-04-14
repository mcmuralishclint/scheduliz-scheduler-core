package main

import (
	"fmt"
	"scheduler-service/cmd/initializer"
)

func main() {
	fmt.Println("Started Scheduler Service")
}

func init() {
	initializer.DbInit()
}
