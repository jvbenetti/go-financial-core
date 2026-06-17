package model

import "time"

type ChargeStatus string

const (
	ChargeStatusPending  ChargeStatus = "pending"
	ChargeStatusPaid     ChargeStatus = "paid"
	ChargeStatusExpired  ChargeStatus = "expired"
	ChargeStatusRefunded ChargeStatus = "refunded"
)

// The Charge it's the invoice that exists before the transaction
type Charge struct {
	ID        string       `gorm:"primaryKey;type:uuid" json:"id"`
	AccountID string       `gorm:"type:uuid;not null;index" json:"account_id"`
	Account   Account      `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"account"`
	Amount    int64        `gorm:"not null" json:"amount"`
	Status    ChargeStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	PixQRCode string       `gorm:"type:text" json:"pix_qr_code"`
	ExpiresAt time.Time    `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
}
