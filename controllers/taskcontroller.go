package controllers

import (
	"golang-api5/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskRequest struct {
	AssignedTo string `json:"assigned_to"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

type UpdateTaskRequest struct {
	AssignedTo string `json:"assigned_to"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var input CreateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	task := models.Task{
		AssignedTo: input.AssignedTo,
		Task:       input.Task,
		Deadline:   deadline,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func FindTask(c *gin.Context) {
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var task models.Task

	// get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input UpdateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedRequest models.Task
	updatedRequest.Deadline = deadline
	updatedRequest.AssignedTo = input.AssignedTo
	updatedRequest.Task = input.Task

	db.Model(&task).Updates(updatedRequest)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
