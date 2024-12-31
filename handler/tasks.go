package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/todos-api/jovi345/formatter"
	"github.com/todos-api/jovi345/task"
)

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService}
}

func (h *taskHandler) AddNewTask(c *gin.Context) {
	var input task.TaskInput

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	clPtr := claims.(*jwt.MapClaims)
	email := (*clPtr)["email"].(string)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newTask, err := h.taskService.AddNewTask(input, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := formatter.SendResponse("success", newTask)
	c.JSON(http.StatusOK, msg)
}

func (h *taskHandler) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskService.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := formatter.SendResponse("success", task)
	c.JSON(http.StatusOK, msg)
}

func (h *taskHandler) GetAllTasks(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, "unauthorized")
		return
	}
	clPtr := (claims).(*jwt.MapClaims)
	email := (*clPtr)["email"].(string)

	tasks, err := h.taskService.GetAllTasks(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := formatter.SendResponse("success", tasks)
	c.JSON(http.StatusOK, msg)
}

func (h *taskHandler) UpdateJobStatus(c *gin.Context) {
	id := c.Param("id")

	foundTask, err := h.taskService.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	status := foundTask.IsCompleted

	data := task.Task{
		ID:          id,
		Job:         foundTask.Job,
		IsCompleted: !status,
		UpdatedAt:   time.Now(),
	}

	task, err := h.taskService.UpdateTodo(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := formatter.SendResponse("success", task)
	c.JSON(http.StatusOK, msg)
}

func (h *taskHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	foundTask, err := h.taskService.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	deletedTask, err := h.taskService.DeleteById(foundTask.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := formatter.SendResponse("success", deletedTask)
	c.JSON(http.StatusOK, msg)
}

func (h *taskHandler) UpdateJob(c *gin.Context) {
	id := c.Param("id")

	var input task.TaskInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.taskService.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data := task.Task{
		ID:        id,
		Job:       input.Job,
		UpdatedAt: time.Now(),
	}

	task, err := h.taskService.UpdateTodo(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	msg := formatter.SendResponse("success", task)
	c.JSON(http.StatusOK, msg)
}
