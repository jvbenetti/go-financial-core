package model

import "time"

// The IdempotencyKey focuses on security and individuality of each collection
type IdempotencyKey struct {
	Key          string    `gorm:"primaryKey;type:varchar(255)" json:"key"`
	ResponseBody []byte    `gorm:"type:bytea" json:"response_body"`
	StatusCode   int       `gorm:"not null" json:"status_code"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
