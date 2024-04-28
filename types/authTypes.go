package types

type AuthService interface {
	Login(params LoginDTO) (*LoginResponse, error)
}

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type UserAuthDetails struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive"`
}

type AuthRepository interface {
	GetUserDetails(username string) (*UserAuthDetails, error)
}
