package main

import (
	"context"
	"log"
	"os"

	"github.com/eu-micaeu/Base/backend/go/database"
	"github.com/eu-micaeu/Base/backend/go/routes"
	"github.com/gin-gonic/gin"
)

// main é o ponto de entrada da aplicação
func main() {
	ctx := context.Background()
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode)
	}
	db, err := database.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	r := routes.Router(db)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("Server listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}