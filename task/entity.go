package task

import "time"

type Task struct {
	ID          string    `json:"id"`
	Job         string    `json:"job"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	Email       string    `json:"email"`
}

type TaskInput struct {
	Job string `json:"job"`
}
