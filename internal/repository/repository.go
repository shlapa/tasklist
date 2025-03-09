package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"tasklist/internal/model"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

// CreateTask создает новую задачу
func (r *TaskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	query := `
		INSERT INTO tasks (title, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	return r.pool.QueryRow(ctx, query, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

// GetAllTasks возвращает все задачи из базы данных
func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	query := `
		SELECT id, title, description, status, created_at, updated_at 
		FROM tasks
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateTask обновляет задачу по ID
func (r *TaskRepository) UpdateTask(ctx context.Context, id int, task *model.Task) error {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`
	return r.pool.QueryRow(ctx, query, task.Title, task.Description, task.Status, id).
		Scan(&task.UpdatedAt)
}

// DeleteTask удаляет задачу по ID
func (r *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
