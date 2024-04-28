package users

import (
	"fmt"

	"github.com/2marks/csts/types"
)

type UserService struct {
	userRepository types.UserRepository
}

func NewService(repo types.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

func (s *UserService) Create(params *types.CreateUserDTO) error {
	userWithEmail, _ := s.userRepository.GetByEmail(params.Email)
	if userWithEmail != nil {
		return fmt.Errorf("user with email %s already exists", params.Email)
	}

	userWithUsername, _ := s.userRepository.GetByUsername(params.Username)
	if userWithUsername != nil {
		return fmt.Errorf("username %s already exists", params.Username)
	}

	return s.userRepository.Create(*params)
}

func (s *UserService) GetAll(params *types.GetAllUsersDTO) (*[]types.User, error) {
	return s.userRepository.GetAll(*params)
}

func (s *UserService) GetById(params *types.GetOneUserDTO) (*types.User, error) {
	user, err := s.userRepository.GetById(params.Id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Activate(params *types.ActivateUserDTO) error {
	user, err := s.userRepository.GetById(params.Id)

	if err != nil {
		return err
	}

	if user.IsActive {
		return fmt.Errorf("user is already active")
	}

	return s.userRepository.Activate(user.Id)
}

func (s *UserService) Deactivate(params *types.DeactivateUserDTO) error {
	user, err := s.userRepository.GetById(params.Id)

	if err != nil {
		return err
	}

	if !user.IsActive {
		return fmt.Errorf("user is already deactivated")
	}

	return s.userRepository.Deactivate(user.Id)
}
