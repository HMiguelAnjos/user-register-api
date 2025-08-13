package domain

import (
	"errors"
	"time"
)

// User is the core Domain entity. It has no framework or storage concerns.
type User struct {
	ID          string
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	ErrEmptyTitle = errors.New("title cannot be empty")
)

// NewUser is a Factory Method that enforces invariants.
func NewUser(id string, title string, description string, now time.Time) (*User, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	return &User{
		ID:          id,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update mutates a User respecting invariants.
func (t *User) Update(title string, description string, done bool, now time.Time) error {
	if title == "" {
		return ErrEmptyTitle
	}
	t.Title = title
	t.Description = description
	t.Done = done
	t.UpdatedAt = now
	return nil
}
