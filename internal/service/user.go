package service

import (
	"errors"

	"github.com/jvbenetti/go-financial-core/internal/dto/request"
	"github.com/jvbenetti/go-financial-core/internal/dto/response"
	"github.com/jvbenetti/go-financial-core/internal/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateUserWithAccount(req *request.UserRequest) (*response.UserResponse, error) {
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

	// If it has error or panic, rollback all!
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
})
}
