package api

import (
	"net/http"
	"scheduler-service/model"
	"scheduler-service/repo"

	"github.com/gin-gonic/gin"
)

func ListSchedules(c *gin.Context, store *repo.MongoStore) {
	tenants, err := store.ListSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something Went Wrong!"})
		return
	}

	c.JSON(http.StatusOK, tenants)
}

func DeleteSchedule(c *gin.Context, store *repo.MongoStore) {
	id := c.Param("id")
	if err := store.DeleteSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tenant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tenant deleted successfully"})
}

func UpdateSchedules(c *gin.Context, store *repo.MongoStore) {
	id := c.Param("id")
	var tenant model.Schedule
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	tenant.ID = id
	if err := store.UpdateSchedule(tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tenant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tenant updated successfully"})
}

func GetTenant(c *gin.Context, store *repo.MongoStore) {
	id := c.Param("id")
	tenant, err := store.GetSchedule(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

func AddTenant(c *gin.Context, store *repo.MongoStore) {
	var tenant model.Schedule
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := store.AddSchedule(tenant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tenant"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tenant added successfully"})
}
