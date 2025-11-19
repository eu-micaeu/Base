package routes

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

    "github.com/eu-micaeu/Base/backend/go/database"
    "github.com/eu-micaeu/Base/backend/go/handlers"
)

// Router é a função que configura as rotas da aplicação
func Router(db *database.DB) *chi.Mux {
    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    uh := handlers.NewUserHandler(db)
    r.Route("/users", func(r chi.Router) {
        r.Get("/", uh.List)
        r.Post("/", uh.Create)
        r.Get("/{id}", uh.Get)
    })

    return r
}
