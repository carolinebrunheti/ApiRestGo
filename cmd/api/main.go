package main

import (
	"fmt"
	"net/http"
	"meu-projeto/internal/handlers"
	"meu-projeto/internal/repository"

	"github.com/go-chi/chi/v5"
	
)

func main() {
    db, _ := repository.NewDB() // A conexão que criamos anteriormente
    r := chi.NewRouter()

    // Rotas do CRUD
    r.Post("/clientes", handlers.CriarCliente(db))
    r.Get("/clientes", handlers.ListarClientes(db))
    r.Put("/clientes/{id}", handlers.AtualizarCliente(db))
    r.Delete("/clientes/{id}", handlers.DeletarCliente(db))

    http.ListenAndServe(":8080", r)
}