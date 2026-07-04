package service

import (
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

	if err := user.Ha
}
