package router

import (
	"rare0b/gin-todo-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	taskHandler := controllers.TaskHandler{
		Db: db,
	}

	r := gin.Default()

	r.GET("/todo", taskHandler.GetAll)
	r.POST("/todo", taskHandler.CreateTask)
	r.PUT("/todo/edit/:id", taskHandler.UpdateTask)
	r.DELETE("/todo/delete/:id", taskHandler.DeleteTask)

	r.Run(":8000")
}
