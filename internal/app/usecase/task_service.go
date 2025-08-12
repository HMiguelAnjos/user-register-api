package app

import (
	"errors"
	"time"
	"userregisterapi/internal/app/ports"
	"userregisterapi/internal/domain"
)

// TaskService is the Application Service (Use Cases).
type TaskService struct {
	repo  ports.TaskRepository
	ids   ports.IDGenerator
	clock func() time.Time
	log   ports.Logger
}

func NewTaskService(repo ports.TaskRepository, ids ports.IDGenerator, log ports.Logger) *TaskService {
	return &TaskService{
		repo:  repo,
		ids:   ids,
		log:   log,
		clock: time.Now,
	}
}

func (s *TaskService) Create(title, description string) (*domain.Task, error) {
	id, err := s.ids.NewID()
	if err != nil {
		return nil, err
	}
	t, err := domain.NewTask(id, title, description, s.clock())
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(t); err != nil {
		return nil, err
	}
	s.log.Info("task_created", map[string]any{"id": t.ID})
	return t, nil
}

func (s *TaskService) Get(id string) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) List() ([]*domain.Task, error) {
	return s.repo.List()
}

func (s *TaskService) Update(id, title, description string, done bool) (*domain.Task, error) {
	t, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, errors.New("task not found")
	}
	if err := t.Update(title, description, done, s.clock()); err != nil {
		return nil, err
	}
	if err := s.repo.Update(t); err != nil {
		return nil, err
	}
	s.log.Info("task_updated", map[string]any{"id": t.ID})
	return t, nil
}

func (s *TaskService) Delete(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	s.log.Info("task_deleted", map[string]any{"id": id})
	return nil
}
