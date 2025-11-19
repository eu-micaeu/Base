package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
    "time"

    "github.com/eu-micaeu/Base/backend/go/database"
    "github.com/eu-micaeu/Base/backend/go/models"
    "github.com/go-chi/chi/v5"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
    DB *database.DB
}

func NewUserHandler(db *database.DB) *UserHandler {
    return &UserHandler{DB: db}
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    err := h.DB.FindAll(r.Context(), "users", bson.D{}, &users)
    if err != nil {
        http.Error(w, "failed to list users", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }
    payload.Name = strings.TrimSpace(payload.Name)
    payload.Email = strings.TrimSpace(strings.ToLower(payload.Email))
    if payload.Name == "" || payload.Email == "" || payload.Password == "" {
        http.Error(w, "name, email and password are required", http.StatusBadRequest)
        return
    }
    if len(payload.Password) < 6 {
        http.Error(w, "password must be at least 6 chars", http.StatusBadRequest)
        return
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "failed to hash password", http.StatusInternalServerError)
        return
    }
    u := models.User{
        Name:         payload.Name,
        Email:        payload.Email,
        PasswordHash: string(hash),
    }
    if u.ID.IsZero() {
        u.ID = primitive.NewObjectID()
    }
    if u.CreatedAt.IsZero() {
        u.CreatedAt = time.Now().UTC()
    }
    _, err = h.DB.InsertOne(r.Context(), "users", u)
    if err != nil {
        http.Error(w, "failed to create user", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
    idHex := chi.URLParam(r, "id")
    var u models.User
    ok, err := h.DB.FindByID(r.Context(), "users", idHex, &u)
    if err != nil {
        http.Error(w, "failed to get user", http.StatusInternalServerError)
        return
    }
    if !ok {
        http.NotFound(w, r)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}
