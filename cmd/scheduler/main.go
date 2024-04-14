package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"scheduler-service/api"
	"scheduler-service/cmd/initializer"
	"scheduler-service/config"
)

var cfg *config.Config

func main() {
	cfg = config.New()
	initalizedConfigs := initializer.DbInit(cfg)
	fmt.Println("Started Scheduler Service")
	router := gin.Default()
	router.GET("/schedules", func(c *gin.Context) { api.ListSchedules(c, initalizedConfigs.MongoStore) })
	router.POST("/schedule", func(c *gin.Context) {})
	router.GET("/schedules/:id", func(c *gin.Context) {})
	router.PUT("/schedules/:id", func(c *gin.Context) {})
	router.DELETE("/schedules/:id", func(c *gin.Context) {})
	router.Run(cfg.ServerPort)
}
