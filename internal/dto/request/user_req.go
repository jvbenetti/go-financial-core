package request

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
}
