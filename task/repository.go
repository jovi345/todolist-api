package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(task Task) (Task, error)
	FindByID(id string) (Task, error)
	FindAll(email string) ([]Task, error)
	Update(data Task) (Task, error)
	DeleteById(task Task) (Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *repository) FindByID(id string) (Task, error) {
	var task Task

	err := r.db.Where("id = ?", id).Find(&task).Error
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *repository) FindAll(email string) ([]Task, error) {
	var tasks []Task

	err := r.db.Where("email = ?", email).Order("updated_at DESC").Find(&tasks).Error
	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func (r *repository) Update(data Task) (Task, error) {
	err := r.db.Where("id = ?", data.ID).Select("job", "is_completed", "updated_at").Updates(data).Error
	if err != nil {
		return Task{}, nil
	}

	return data, nil
}

func (r *repository) DeleteById(task Task) (Task, error) {
	err := r.db.Where("id = ?", task.ID).Delete(&task).Error
	if err != nil {
		return Task{}, err
	}

	return task, nil
}
