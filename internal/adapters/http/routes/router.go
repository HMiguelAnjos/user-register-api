package routers

import (
	"net/http"

	"userregisterapi/internal/adapters/http/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(c *controllers.UserController) http.Handler {
	r := chi.NewRouter()

	// Middlewares úteis (id de request, recovery, compressão, etc.)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			// collection
			r.Get("/", c.ListUsers)
			r.Post("/", c.CreateUser)

			// item
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, req *http.Request) {
					c.GetUser(w, req, chi.URLParam(req, "id"))
				})
				r.Put("/", func(w http.ResponseWriter, req *http.Request) {
					c.UpdateUser(w, req, chi.URLParam(req, "id"))
				})
				r.Delete("/", func(w http.ResponseWriter, req *http.Request) {
					c.DeleteUser(w, req, chi.URLParam(req, "id"))
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
