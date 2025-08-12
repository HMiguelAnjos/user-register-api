package postgres

import (
	"database/sql"
	"userregisterapi/internal/app/ports"
	"userregisterapi/internal/common"
	"userregisterapi/internal/domain"
)

var _ ports.TaskRepository = (*TaskRepoPostgres)(nil)

type TaskRepoPostgres struct {
	db *sql.DB
}

func NewTaskRepoPostgres(db *sql.DB) *TaskRepoPostgres {
	return &TaskRepoPostgres{db: db}
}

func (r *TaskRepoPostgres) Save(task *domain.Task) error {
	_, err := r.db.Exec(`INSERT INTO tasks (id, title, description, done, created_at, updated_at)
                         VALUES ($1, $2, $3, $4, $5, $6)`,
		task.ID, task.Title, task.Description, task.Done, task.CreatedAt, task.UpdatedAt)
	return err
}

func (r *TaskRepoPostgres) GetByID(id string) (*domain.Task, error) {
	row := r.db.QueryRow(`SELECT id, title, description, done, created_at, updated_at FROM tasks WHERE id=$1`, id)
	t := &domain.Task{}
	if err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt, &t.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

func (r *TaskRepoPostgres) List() ([]*domain.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, description, done, created_at, updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*domain.Task
	for rows.Next() {
		t := &domain.Task{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepoPostgres) Update(task *domain.Task) error {
	res, err := r.db.Exec(`UPDATE tasks SET title=$1, description=$2, done=$3, updated_at=$4 WHERE id=$5`,
		task.Title, task.Description, task.Done, task.UpdatedAt, task.ID)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return common.ErrNotFound
	}
	return nil
}

func (r *TaskRepoPostgres) Delete(id string) error {
	res, err := r.db.Exec(`DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return common.ErrNotFound
	}
	return nil
}
