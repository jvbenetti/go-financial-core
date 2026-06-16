package model

import "time"

// The Transaction is macro event (pix to: XYZ)
type Transaction struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// The LedgerEntry is which affected line
type LedgerEntry struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	AccountID     string    `json:"account_id"`
	Amount        int64     `json:"amount"` // Negativo se for Débito, Positivo se for Crédito
	CreatedAt     time.Time `json:"created_at"`
}
