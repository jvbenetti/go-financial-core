package model

import "time"

type Account struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   int64     `json:"balance"`  // Balance in cents
	Currency  string    `json:"currency"` // Ex: "BRL"
	Status    string    `json:"status"`   // "active", "blocked"
	UpdatedAt time.Time `json:"updated_at"`
}
