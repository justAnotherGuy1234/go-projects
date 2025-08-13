package dto

type SignUpUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
