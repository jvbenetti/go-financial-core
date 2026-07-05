package response

import (
	"time"
)

type AccountResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
