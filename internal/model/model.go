package model

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	StatusNew        TaskStatus = "new"         // Новая
	StatusInProgress TaskStatus = "in_progress" // В процессе
	StatusDone       TaskStatus = "done"        // Выполненна
)

type Task struct {
	ID          int        `json:"id"`          // Уникальный идентификатор задачи
	Title       string     `json:"title"`       // Название задачи
	Description string     `json:"description"` // Описание задачи
	Status      TaskStatus `json:"status"`      // Статус задачи (new, in_progress, done)
	CreatedAt   time.Time  `json:"created_at"`  // Время создания задачи
	UpdatedAt   time.Time  `json:"updated_at"`  // Время последнего обновления задачи
}

func (s TaskStatus) Validate() error {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return nil
	default:
		return errors.New("invalid status")
	}
}
