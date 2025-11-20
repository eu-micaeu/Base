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
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(cors.New(cors.Config{
        // Permite qualquer origem e adiciona cabeçalho corretamente sem precisar de redirect.
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposeHeaders:    []string{"Link"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }))

    r.GET("/health", func(c *gin.Context) {
        c.String(200, "ok")
    })

    uh := handlers.NewUserHandler(db)

    auth := r.Group("/auth")
    {
        auth.Use(middlewares.CORSMiddleware())
        auth.POST("/register", uh.Register)
        auth.POST("/login", uh.Login)
    }

    users := r.Group("/users")
    {
        // Aplica middleware do pacote middlewares ao grupo /users
        users.Use(middlewares.CORSMiddleware())
        // Suporte a listagem e busca por id sem barra final
        users.GET("", uh.List)
        users.GET(":id", uh.Get)
        users.POST("", uh.Create)
    }

    return r
}
