package service

import (
	"errors"

	"github.com/jvbenetti/go-financial-core/internal/dto/request"
	"github.com/jvbenetti/go-financial-core/internal/dto/response"
	"github.com/jvbenetti/go-financial-core/internal/model"
	"golang.org/x/text/currency"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateUserWithAccount(req *request.UserRequest) (*response.UserResponse, error) {
	// 1: Creating the user struct
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Document: req.Document,
		Phone:    req.Phone,
		Role:     model.UserRoleCustomer,
	}

	if err := user.HashPassword(req.Password); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Init transaction
	tx := s.DB.Begin()

	// 2: If it has error or panic, rollback all!
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 3: Save the user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 4: Save the account to the new user
	account := model.Account{
		UserID: user.ID,
		Balance: 0,
		Currency: model.CurrencyBRL,
		Status: model.AccountStatusActive,
	}
	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback() // If its has a error, rollback
		return nil, errors.New("falha ao criar a conta bancária")
	}

})
}
