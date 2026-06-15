package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"` // CPF ou CNPJ
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
