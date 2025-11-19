package main

import (
	"context"
	"log"
	"net/http"

	"github.com/eu-micaeu/Base/backend/go/database"
	"github.com/eu-micaeu/Base/backend/go/routes"
)

// main é o ponto de entrada da aplicação
func main() {
	ctx := context.Background()
	db, err := database.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	r := routes.Router(db)
	addr := ":8080"
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}