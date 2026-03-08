//@title: task API
//@version: 1.0
//@description: esta e uma API para gerenciar tarefas
//@host: localhost:8080
//@BasePath: /

package main

import (
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary: create task
// @Description: create a new task
// @Tags: tasks
// @Accept: json
// @Produce: json
// @Param: task body models.Task true "task to create"
// @Success: 201 {object} map[string]interface{}
// @Failure: 400 {object} map[string]string
// @Failure: 500 {object} map[string]string
// @Router: /tasks [post]

// @Summary: list tasks
// @Description: list all tasks
// @Tags: tasks
// @Accept: json
// @Produce: json
// @Param: id path string true "task id"
// @Param: task body models.Task true "task to update"
// @Success: 200 {array} models.Task
// @Failure: 400 {object} map[string]string
// @failure: 404 {object} map[string]string
// @Failure: 500 {object} map[string]string
// @Router: /tasks [get]

// @Summary: update task
// @Description: update a task by id
// @Tags: tasks
// @Accept: json
// @Produce: json
// @Param: id path string true "task id"
// @Param: task body models.Task true "task to update"
// @Success: 200 {object} map[string]string
// @Failure: 400 {object} map[string]string
// @failure: 404 {object} map[string]string
// @Failure: 500 {object} map[string]string
// @Router: /tasks/{id} [put]

// @Summary: delete task
// @Description: delete a task by id
// @Tags: tasks
// @Accept: json
// @Produce: json
// @Param: id path string true "task id"
// @Success: 200 {object} map[string]string
// @Failure: 400 {object} map[string]string
// @failure: 404 {object} map[string]string
// @Failure: 500 {object} map[string]string
// @Router: /tasks/{id} [delete]

func main() {

	router := gin.Default()

	router.Use(cors.new(cors.config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tasks", listTasks)

	router.POST("/tasks", createtask)

	router.PUT("/tasks/:id", updateTask)

	router.DELETE("/tasks/:id", deleteTask)

	router.Run(":8080")
}
