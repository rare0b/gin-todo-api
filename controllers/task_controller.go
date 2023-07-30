package controllers

import (
	"errors"
	"net/http"
	"rare0b/gin-todo-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Db *gorm.DB
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	var tasks []models.Task

	result := h.Db.Find(&tasks)

	if result.Error != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	task := models.Task{}
	err := c.BindJSON(&task)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
	} else {
		h.Db.Create(&task)
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	task := models.Task{}
	id := c.Param("id")

	// データの検索に失敗した場合にErrRecordNotFoundを返します
	result := h.Db.First(&task, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	// 指定のidのtaskのtitleを、受け取ったJSONのtitleに変更
	jsonTask := models.Task{}
	err := c.BindJSON(&jsonTask)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
	} else {
		task.Title = jsonTask.Title
		h.Db.Save(&task)
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	task := models.Task{}
	id := c.Param("id")

	result := h.Db.Delete(&task, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
	} else if result.Error != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
