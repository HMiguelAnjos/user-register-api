package httpadapter

import (
	"encoding/json"
	"net/http"
	"time"

	"userregisterapi/internal/app/dto"
	app "userregisterapi/internal/app/usecase"
	"userregisterapi/internal/common"
	"userregisterapi/internal/domain"
)

// TaskController is a Controller that translates HTTP <-> UseCases (Application layer).
type TaskController struct {
	svc *app.TaskService
}

func NewTaskController(svc *app.TaskService) *TaskController {
	return &TaskController{svc: svc}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	t, err := c.svc.Create(req.Title, req.Description)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, toTaskResponse(t))
}

func (c *TaskController) GetTask(w http.ResponseWriter, r *http.Request, id string) {
	t, err := c.svc.Get(id)
	if err != nil {
		if err == common.ErrNotFound {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, toTaskResponse(t))
}

func (c *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	list, err := c.svc.List()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	out := make([]dto.TaskResponse, 0, len(list))
	for _, t := range list {
		out = append(out, toTaskResponse(t))
	}
	writeJSON(w, http.StatusOK, out)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request, id string) {
	var req dto.UpdateTaskRequest
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
	writeJSON(w, http.StatusOK, toTaskResponse(t))
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request, id string) {
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

func toTaskResponse(t *domain.Task) dto.TaskResponse {
	return dto.TaskResponse{
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
