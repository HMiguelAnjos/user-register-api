package domain

import (
    "errors"
    "time"
)

// Task is the core Domain entity. It has no framework or storage concerns.
type Task struct {
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

// NewTask is a Factory Method that enforces invariants.
func NewTask(id string, title string, description string, now time.Time) (*Task, error) {
    if title == "" {
        return nil, ErrEmptyTitle
    }
    return &Task{
        ID:          id,
        Title:       title,
        Description: description,
        Done:        false,
        CreatedAt:   now,
        UpdatedAt:   now,
    }, nil
}

// Update mutates a Task respecting invariants.
func (t *Task) Update(title string, description string, done bool, now time.Time) error {
    if title == "" {
        return ErrEmptyTitle
    }
    t.Title = title
    t.Description = description
    t.Done = done
    t.UpdatedAt = now
    return nil
}
