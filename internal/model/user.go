package model

import "time"

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

// User is about CPF or CNPJ person
type User struct {
	ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"user_name"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Document  string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"document"` // Unique CPF or CNPJ
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone     int64     `gorm:"type:varchar()"`
	Role      UserRole  `gorm:"default:'customer';type:varchar(50)" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
