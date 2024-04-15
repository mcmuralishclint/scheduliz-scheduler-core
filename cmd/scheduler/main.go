package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"scheduler-service/api"
	"scheduler-service/cmd/initializer"
	"scheduler-service/config"
	"scheduler-service/domain/scheduler"
)

var cfg *config.Config

func main() {
	cfg = config.New()
	initalizedConfigs := initializer.DbInit(cfg)
	fmt.Println("Started Scheduler Service")
	router := gin.Default()
	router.GET("/schedules", func(c *gin.Context) { api.ListSchedules(c, initalizedConfigs.MongoStore) })
	router.POST("/schedule", func(c *gin.Context) { api.AddSchedule(c, initalizedConfigs.MongoStore) })
	router.GET("/schedules/:id", func(c *gin.Context) { api.GetSchedule(c, initalizedConfigs.MongoStore) })
	router.PUT("/schedules/:id", func(c *gin.Context) { api.UpdateSchedule(c, initalizedConfigs.MongoStore) })
	router.DELETE("/schedules/:id", func(c *gin.Context) { api.DeleteSchedule(c, initalizedConfigs.MongoStore) })
	go func() {
		if err := scheduler.ScheduleJobs(context.Background()); err != nil {
			log.Fatalf("Error scheduling jobs: %v", err)
		}
	}()
	router.Run(cfg.ServerPort)
}
