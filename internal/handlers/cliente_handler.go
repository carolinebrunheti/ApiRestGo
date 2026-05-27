package handlers

import(
	"database/sql"
    "encoding/json"
    "net/http"
    "meu-projeto/internal/models"
    "github.com/go-chi/chi/v5"
)

//Exemplo de post
func CriarCliente(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var c models.Cliente
        json.NewDecoder(r.Body).Decode(&c)

        query := "INSERT INTO clientes (nome, email) VALUES ($1, $2) RETURNING id"
        err := db.QueryRow(query, c.Nome, c.Email).Scan(&c.ID)
        
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(c)
    }
}