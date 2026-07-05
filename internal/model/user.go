package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

// User is about CPF or CNPJ person
type User struct {
	ID        uint      `gorm:"primaryKey;type:uuid" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"user_name"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Document  string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"document"` // Unique CPF or CNPJ
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Phone     string    `gorm:"type:varchar()"`
	Role      UserRole  `gorm:"default:'customer';type:varchar(50)" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (u *User) HashPassword(plainPassword string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14) // 14 é o "cost" (security level)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil
}
