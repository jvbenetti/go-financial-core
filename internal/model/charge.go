package model

import "time"

type ChargeStatus string

const (
	pending  ChargeStatus = "pending"
	paid     ChargeStatus = "paid"
	expired  ChargeStatus = "expired"
	refunded ChargeStatus = "refunded"
)

type Charge struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"` // Conta que vai receber o dinheiro
	Account   Account   `json:"account"`
	Amount    int64     `json:"amount"`      // Valor em centavos
	Status    string    `json:"status"`      // "pending", "paid", "expired", "refunded"
	PixQRCode string    `json:"pix_qr_code"` // Dados do QR Code simulado
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
