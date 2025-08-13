package app

import (
	"errors"
	"time"
	"userregisterapi/internal/app/ports"
	"userregisterapi/internal/domain"
)

// UserService is the Application Service (Use Cases).
type UserService struct {
	repo  ports.UserRepository
	ids   ports.IDGenerator
	clock func() time.Time
	log   ports.Logger
}

func NewUserService(repo ports.UserRepository, ids ports.IDGenerator, log ports.Logger) *UserService {
	return &UserService{
		repo:  repo,
		ids:   ids,
		log:   log,
		clock: time.Now,
	}
}

func (s *UserService) Create(title, description string) (*domain.User, error) {
	id, err := s.ids.NewID()
	if err != nil {
		return nil, err
	}
	t, err := domain.NewUser(id, title, description, s.clock())
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(t); err != nil {
		return nil, err
	}
	s.log.Info("User_created", map[string]any{"id": t.ID})
	return t, nil
}

func (s *UserService) Get(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) List() ([]*domain.User, error) {
	return s.repo.List()
}

func (s *UserService) Update(id, title, description string, done bool) (*domain.User, error) {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, errors.New("User not found")
	}
	if err := t.Update(title, description, done, s.clock()); err != nil {
		return nil, err
	}
	if err := s.repo.Update(t); err != nil {
		return nil, err
	}
	s.log.Info("User_updated", map[string]any{"id": t.ID})
	return t, nil
}

func (s *UserService) Delete(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	s.log.Info("User_deleted", map[string]any{"id": id})
	return nil
}
