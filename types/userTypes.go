package types

import "time"

type roleTypes struct {
	agent    string
	customer string
	admin    string
}

var UserRole = roleTypes{
	agent:    "agent",
	customer: "customer",
	admin:    "admin",
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3"`
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=3"`
	Role     string `json:"role" validate:"required"`
}

type GetAllUsersDTO struct {
	Page        int    `json:"page" validate:"required"`
	PerPage     int    `json:"perPage" validate:"required"`
	SearchQuery string `json:"searchQuery"`
}

type GetOneUserDTO struct {
	Id int `json:"id"`
}

type ActivateUserDTO struct {
	Id int `json:"id"`
}

type DeactivateUserDTO struct {
	Id int `json:"id"`
}

type UserRepository interface {
	Create(CreateUserDTO) error
	GetById(int) (*User, error)
	GetByEmail(string) (*User, error)
	GetByUsername(string) (*User, error)
	GetAll(GetAllUsersDTO) (*[]User, error)
	Activate(int) error
	Deactivate(int) error
}

type UserService interface {
	Create(*CreateUserDTO) error
	GetAll(*GetAllUsersDTO) (*[]User, error)
	Activate(*ActivateUserDTO) error
	Deactivate(*DeactivateUserDTO) error
	GetById(*GetOneUserDTO) (*User, error)
}
