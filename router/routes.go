package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/todos-api/jovi345/config"
	"github.com/todos-api/jovi345/handler"
	"github.com/todos-api/jovi345/middleware"
	"github.com/todos-api/jovi345/task"
	"github.com/todos-api/jovi345/user"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()
	db, _ := config.InitDB()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://frontend-dot-tugas-15-31122024.et.r.appspot.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	api := r.Group("/api/v1")

	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.LoginUser)
	api.POST("/user/refreshtoken", userHandler.RefreshToken)
	api.DELETE("/user/logout", userHandler.Logout)

	api.GET("/task/get/:id", middleware.VerifyToken(), taskHandler.GetTaskById)
	api.GET("/task/getall", middleware.VerifyToken(), taskHandler.GetAllTasks)
	api.POST("/task/add", middleware.VerifyToken(), taskHandler.AddNewTask)
	api.PATCH("/task/edit/status/:id", middleware.VerifyToken(), taskHandler.UpdateJobStatus)
	api.PATCH("/task/edit/job/:id", middleware.VerifyToken(), taskHandler.UpdateJob)
	api.DELETE("/task/delete/:id", middleware.VerifyToken(), taskHandler.DeleteById)

	return r
}
