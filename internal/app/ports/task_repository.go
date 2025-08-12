package ports

import "userregisterapi/internal/domain"

// TaskRepository is the Repository pattern port.
type TaskRepository interface {
	Save(task *domain.Task) error
	GetByID(id string) (*domain.Task, error)
	List() ([]*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id string) error
}
