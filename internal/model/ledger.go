package model

import "time"

// The Transaction is macro event (pix to: XYZ)
type Transaction struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
