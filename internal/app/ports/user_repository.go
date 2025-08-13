package ports

import "userregisterapi/internal/domain"

// UserRepository is the Repository pattern port.
type UserRepository interface {
	Save(users *domain.User) error
	GetByID(id string) (*domain.User, error)
	List() ([]*domain.User, error)
	Update(users *domain.User) error
	Delete(id string) error
}
