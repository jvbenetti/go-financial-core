package response

import "github.com/jvbenetti/go-financial-core/internal/model"

type UserResponse struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Document string         `json:"document"`
	Phone    string         `json:"phone"`
	Role     model.UserRole `json:"role"`
	// Never return the Password
}
