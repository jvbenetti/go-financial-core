package model

import "time"

type CurrencyCode string

const (
	CurrencyBRL CurrencyCode = "BRL"
	CurrencyUSD CurrencyCode = "USD"
	CurrencyEUR CurrencyCode = "EUR"
)

type AccountStatus string

const (
	AccountStatusPending AccountStatus = "pending"
	AccountStatusActive  AccountStatus = "active"
	AccountStatusBlocked AccountStatus = "blocked"
)

// Account involves everything about transactions and Charges
type Account struct {
	ID        string        `gorm:"primaryKey;type:uuid" json:"id"`
	UserID    string        `gorm:"type:uuid;not null;index" json:"user_id"`
	User      User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user"`
	Balance   int64         `gorm:"not null;default:0" json:"balance"`
	Currency  CurrencyCode  `gorm:"type:varchar(3);not null;default:'BRL'" json:"currency"`
	Status    AccountStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
