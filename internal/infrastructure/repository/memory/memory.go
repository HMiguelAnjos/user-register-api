package memory

import (
	"sync"
	"userregisterapi/internal/app/ports"
	"userregisterapi/internal/common"
	"userregisterapi/internal/domain"
)

// Ensure implementation
var _ ports.UserRepository = (*UserRepoMemory)(nil)

type UserRepoMemory struct {
	mu    sync.RWMutex
	items map[string]*domain.User
}

func NewUserRepoMemory() *UserRepoMemory {
	return &UserRepoMemory{items: make(map[string]*domain.User)}
}

func (r *UserRepoMemory) Save(User *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[User.ID] = User
	return nil
}

func (r *UserRepoMemory) GetByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, common.ErrNotFound
	}
	return t, nil
}

func (r *UserRepoMemory) List() ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*domain.User, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	return out, nil
}

func (r *UserRepoMemory) Update(User *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[User.ID]; !ok {
		return common.ErrNotFound
	}
	r.items[User.ID] = User
	return nil
}

func (r *UserRepoMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return common.ErrNotFound
	}
	delete(r.items, id)
	return nil
}
