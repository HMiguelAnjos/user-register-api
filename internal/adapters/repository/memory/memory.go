package memory

import (
	"sync"
	"userregisterapi/internal/app/ports"
	"userregisterapi/internal/common"
	"userregisterapi/internal/domain"
)

// Ensure implementation
var _ ports.TaskRepository = (*TaskRepoMemory)(nil)

type TaskRepoMemory struct {
	mu    sync.RWMutex
	items map[string]*domain.Task
}

func NewTaskRepoMemory() *TaskRepoMemory {
	return &TaskRepoMemory{items: make(map[string]*domain.Task)}
}

func (r *TaskRepoMemory) Save(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[task.ID] = task
	return nil
}

func (r *TaskRepoMemory) GetByID(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, common.ErrNotFound
	}
	return t, nil
}

func (r *TaskRepoMemory) List() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*domain.Task, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	return out, nil
}

func (r *TaskRepoMemory) Update(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[task.ID]; !ok {
		return common.ErrNotFound
	}
	r.items[task.ID] = task
	return nil
}

func (r *TaskRepoMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return common.ErrNotFound
	}
	delete(r.items, id)
	return nil
}
