package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"userregisterapi/internal/app/dto"
	app "userregisterapi/internal/app/usecase"
	"userregisterapi/internal/common"
	"userregisterapi/internal/domain"
)

// UserController is a Controller that translates HTTP <-> UseCases (Application layer).
type UserController struct {
	svc *app.UserService
}

func NewUserController(svc *app.UserService) *UserController {
	return &UserController{svc: svc}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	t, err := c.svc.Create(req.Title, req.Description)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, toUserResponse(t))
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request, id string) {
	t, err := c.svc.Get(id)
	if err != nil {
		if err == common.ErrNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, toUserResponse(t))
}

func (c *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	list, err := c.svc.List()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	out := make([]dto.UserResponse, 0, len(list))
	for _, t := range list {
		out = append(out, toUserResponse(t))
	}
	writeJSON(w, http.StatusOK, out)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request, id string) {
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	t, err := c.svc.Update(id, req.Title, req.Description, req.Done)
	if err != nil {
		if err == common.ErrNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, toUserResponse(t))
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, id string) {
	if err := c.svc.Delete(id); err != nil {
		if err == common.ErrNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- Helpers ---

func writeDomainError(w http.ResponseWriter, err error) {
	switch err {
	case domain.ErrEmptyTitle:
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}

func toUserResponse(t *domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
		CreatedAt:   t.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:   t.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
