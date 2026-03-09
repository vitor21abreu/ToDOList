package delivery

import (
	"go-crud-api/internal/delivery/dependencias"
	"go-crud-api/internal/delivery/handlers"
	"go-crud-api/internal/delivery/middlewares/cors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {

	container := dependencias.Setup()
	router := gin.Default()

	router.Use(cors.new(cors.Connfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	err := container.Invoke(func(taskHandler *handlers.TaskHandler) {
		router.POST("/tasks", taskHandler.CreateTask)
		router.GET("/tasks", taskHandler.GetTask)
		router.PUT("/tasks/:id", taskHandler.UpdateTask)
		router.DELETE("/tasks/:id", taskHandler.DeleteTask)

		log.Println("Starting server on :8080")
		router.Run(":8080")
	})

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
