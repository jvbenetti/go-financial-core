package model

import "time"

// The Transaction is macro event (pix to: XYZ)
// Transaction just keep reasons and descriptions
type Transaction struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// The LedgerEntry it's the immutable mathematical movement
type LedgerEntry struct {
	ID            string      `gorm:"primaryKey;type:uuid" json:"id"`
	TransactionID string      `gorm:"type:uuid;not null;index" json:"transaction_id"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"transaction"`
	AccountID     string      `gorm:"type:uuid;not null;index" json:"account_id"`
	Account       Account     `gorm:"foreignKey:AccountID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"account"`
	Amount        int64       `gorm:"not null" json:"amount"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
}
