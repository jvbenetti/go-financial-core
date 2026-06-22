package request

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
}
