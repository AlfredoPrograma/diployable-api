package dtos

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
