package httpadapter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(c *TaskController) http.Handler {
	r := chi.NewRouter()

	// Middlewares úteis (id de request, recovery, compressão, etc.)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			// collection
			r.Get("/", c.ListTasks)
			r.Post("/", c.CreateTask)

			// item
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, req *http.Request) {
					c.GetTask(w, req, chi.URLParam(req, "id"))
				})
				r.Put("/", func(w http.ResponseWriter, req *http.Request) {
					c.UpdateTask(w, req, chi.URLParam(req, "id"))
				})
				r.Delete("/", func(w http.ResponseWriter, req *http.Request) {
					c.DeleteTask(w, req, chi.URLParam(req, "id"))
				})
			})
		})
	})

	// 404 amigável (opcional)
	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":"not found"}`))
	}))

	return r
}
