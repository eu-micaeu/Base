package routes

import (
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "github.com/eu-micaeu/Base/backend/go/database"
    "github.com/eu-micaeu/Base/backend/go/handlers"
    "github.com/eu-micaeu/Base/backend/go/middlewares"
)

// Router configura e retorna o engine do Gin
func Router(db *database.DB) *gin.Engine {
    r := gin.New()

    uh := handlers.NewUserHandler(db)

    auth := r.Group("/auth")
    {
        auth.Use(middlewares.CORSMiddleware())
        auth.POST("/register", uh.Register)
        auth.POST("/login", uh.Login)
    }

    return r
}
