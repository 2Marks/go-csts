package auth

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/2marks/csts/config"
	"github.com/2marks/csts/types"
	"github.com/2marks/csts/utils"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	authRepository types.AuthRepository
}

func NewService(repo types.AuthRepository) *AuthService {
	return &AuthService{authRepository: repo}
}

func (s *AuthService) Login(params types.LoginDTO) (*types.LoginResponse, error) {
	user, err := s.authRepository.GetUserDetails(params.Username)

	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user is not active")
	}

	if !utils.ValidatePassword(user.Password, params.Password) {
		return nil, fmt.Errorf("login credentials incorrect")
	}

	loginResponse := types.LoginResponse{
		Id:    user.Id,
		Token: generateAuthToken(user.Id),
	}

	return &loginResponse, nil
}

func generateAuthToken(userId int) string {
	expiration := time.Second * time.Duration(config.Envs.JwtExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Envs.JwtSecret))
	if err != nil {
		log.Printf("error occured while generating auth token %v", err)
		return ""
	}

	return tokenString
}
