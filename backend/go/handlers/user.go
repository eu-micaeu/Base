package handlers

import (
    "net/http"
    "strings"
    "time"

    "github.com/eu-micaeu/Base/backend/go/database"
    "github.com/eu-micaeu/Base/backend/go/models"
    "github.com/eu-micaeu/Base/backend/go/utils"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
    DB *database.DB
}

// NewUserHandler cria um novo UserHandler
func NewUserHandler(db *database.DB) *UserHandler {
    return &UserHandler{DB: db}
}

// Create cria um novo usuário
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

// List retorna todos os usuários
func (h *UserHandler) List(c *gin.Context) {
    var users []models.User
    if err := h.DB.FindAll(c.Request.Context(), "users", bson.M{}, &users); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

// Get retorna um usuário por ID
func (h *UserHandler) Get(c *gin.Context) {
    id := c.Param("id")
    var u models.User
    found, err := h.DB.FindByID(c.Request.Context(), "users", id, &u)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, u)
}

// Register cria um novo usuário (similar a Create, mas impede e-mail duplicado)
func (h *UserHandler) Register(c *gin.Context) {
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
    // Verificar e-mail existente
    var existing models.User
    found, err := h.DB.FindOne(c.Request.Context(), "users", bson.M{"email": payload.Email}, &existing)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check existing user"})
        return
    }
    if found {
        c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
        return
    }
    // Criar usuário
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
    c.JSON(http.StatusCreated, gin.H{"message": "user registered", "user": u})
}

// Login valida credenciais e retorna dados básicos do usuário
func (h *UserHandler) Login(c *gin.Context) {
    var payload struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
        return
    }
    email := strings.TrimSpace(strings.ToLower(payload.Email))
    if email == "" || payload.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
        return
    }
    var u models.User
    found, err := h.DB.FindOne(c.Request.Context(), "users", bson.M{"email": email}, &u)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup user"})
        return
    }
    if !found {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(payload.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    // Gerar token JWT
    token, err := utils.GenerateJWT(u.ID.Hex(), u.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }
    
    // Definir cookie JWT
    c.SetCookie(
        "jwt", token,
        60*60*72, // 72 horas
        "/",      // path
        "",       // domain (vazio = atual)
        false,     // secure (true em produção)
        true,      // httpOnly
    )

    c.JSON(http.StatusOK, gin.H{
        "message": "login successful",
        "user": u,
    })
}