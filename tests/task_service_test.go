package tests

import (
	"testing"

	"userregisterapi/internal/adapters/id"
	"userregisterapi/internal/adapters/logger"
	app "userregisterapi/internal/app/usecase"
	"userregisterapi/internal/infrastructure/repository/memory"
)

func TestCreateListTask_Unit(t *testing.T) {
	repo := memory.NewUserRepoMemory()
	ids := id.NewUUIDGenerator()
	lg := logger.NewStdLogger()
	svc := app.NewUserService(repo, ids, lg)

	created, err := svc.Create("Learn Go", "Practice SOLID + Hexagonal")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if created.ID == "" {
		t.Fatalf("expected ID to be generated")
	}

	list, err := svc.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 task, got %d", len(list))
	}
}
