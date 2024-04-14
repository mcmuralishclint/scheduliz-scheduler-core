package main

import (
	"fmt"
	"scheduler-service/cmd/initializer"
)

func main() {
	fmt.Println("Started Worker Service")
}

func init() {
	initializer.DbInit()
}
