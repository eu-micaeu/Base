package routes

import (
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "github.com/eu-micaeu/Base/backend/go/database"
    "github.com/eu-micaeu/Base/backend/go/handlers"
)

// Router configura e retorna o engine do Gin
func Router(db *database.DB) *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
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
    users := r.Group("/users")
    {
        users.GET("/", uh.List)
        users.POST("/", uh.Create)
        users.GET("/:id", uh.Get)
    }

    return r
}
