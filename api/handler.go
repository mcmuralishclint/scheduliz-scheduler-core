package api

import (
	"net/http"
	"scheduler-service/model"
	"scheduler-service/repo"

	"github.com/gin-gonic/gin"
)

func ListSchedules(c *gin.Context, store *repo.MongoStore) {
	schedules, err := store.ListSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something Went Wrong!"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func DeleteSchedule(c *gin.Context, store *repo.MongoStore) {
	id := c.Param("id")
	if err := store.DeleteSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}

func UpdateSchedule(c *gin.Context, store *repo.MongoStore) {
	id := c.Param("id")
	var schedule model.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	schedule.ID = id
	if err := store.UpdateSchedule(schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
}

func GetSchedule(c *gin.Context, store *repo.MongoStore) {
	id := c.Param("id")
	schedule, err := store.GetSchedule(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func AddSchedule(c *gin.Context, store *repo.MongoStore) {
	var schedule model.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := store.AddSchedule(schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add schedule"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule added successfully"})
}
