package task

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	AddNewTask(input TaskInput, enail string) (Task, error)
	GetTaskById(id string) (Task, error)
	GetAllTasks(email string) ([]Task, error)
	UpdateTodo(data Task) (Task, error)
	DeleteById(id string) (Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddNewTask(input TaskInput, email string) (Task, error) {
	taskID := uuid.New().String()
	created_at := time.Now()
	updated_at := created_at

	task := Task{
		ID:          taskID,
		Job:         input.Job,
		IsCompleted: false,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
		Email:       email,
	}

	res, err := s.repository.Save(task)
	if err != nil {
		return Task{}, err
	}

	return res, nil
}

func (s *service) GetTaskById(id string) (Task, error) {
	task, err := s.repository.FindByID(id)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *service) GetAllTasks(email string) ([]Task, error) {
	tasks, err := s.repository.FindAll(email)
	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func (s *service) UpdateTodo(data Task) (Task, error) {
	updatedTask, err := s.repository.Update(data)
	if err != nil {
		return Task{}, err
	}

	return updatedTask, nil
}

func (s *service) DeleteById(id string) (Task, error) {
	task, err := s.repository.FindByID(id)
	if err != nil {
		return Task{}, err
	}

	deletedTask, err := s.repository.DeleteById(task)
	if err != nil {
		return Task{}, err
	}

	return deletedTask, nil
}
