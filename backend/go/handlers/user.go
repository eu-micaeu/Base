package handlers

import (
    "net/http"
    "strings"
    "time"

    "github.com/eu-micaeu/Base/backend/go/database"
    "github.com/eu-micaeu/Base/backend/go/models"
    "github.com/gin-gonic/gin"
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

func (h *UserHandler) List(c *gin.Context) {
    var users []models.User
    if err := h.DB.FindAll(c.Request.Context(), "users", bson.D{}, &users); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Create(c *gin.Context) {
    var payload struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
        return
    }
    payload.Name = strings.TrimSpace(payload.Name)
    payload.Email = strings.TrimSpace(strings.ToLower(payload.Email))
    if payload.Name == "" || payload.Email == "" || payload.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "name, email and password are required"})
        return
    }
    if len(payload.Password) < 6 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 chars"})
        return
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
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
    if _, err = h.DB.InsertOne(c.Request.Context(), "users", u); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
        return
    }
    c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) Get(c *gin.Context) {
    idHex := c.Param("id")
    var u models.User
    ok, err := h.DB.FindByID(c.Request.Context(), "users", idHex, &u)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
        return
    }
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, u)
}
