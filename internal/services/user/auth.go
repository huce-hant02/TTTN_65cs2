package user

import (
	"context"
	"mono-base/internal/repositories"
)

type AuthService interface {
	CreateAuthToken(ctx context.Context, input CreateAuthTokenInput) (*CreateAuthTokenOutput, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) CreateAuthToken(ctx context.Context, input CreateAuthTokenInput) (*CreateAuthTokenOutput, error) {
	// Business logic here
	return &CreateAuthTokenOutput{
		AccessToken:  "hehe",
		RefreshToken: "hihi",
	}, nil
}
