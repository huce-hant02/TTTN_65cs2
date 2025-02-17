package user

import (
	"context"
	"mono-base/internal/repositories"
)

// GetUserByIdInput DTO Input for GetUserById UseCase
type GetUserByIdInput struct {
	ID int `json:"id"`
}

// GetUserByIdOutput DTO Output for GetUserById UseCase
type GetUserByIdOutput struct {
	ID        int     `json:"id"`
	UserName  string  `json:"user_name"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

// GetUserByIdUseCase defines an interface for retrieving a user by their ID.
type GetUserByIdUseCase interface {
	// Execute retrieves a user based on the provided ID.
	// It takes a context and GetUserByIdInput containing the user ID,
	// and returns a GetUserByIdOutput containing user details and an error if the retrieval fails.
	Execute(ctx context.Context, input GetUserByIdInput) (*GetUserByIdOutput, error)
}

type getUserByIdUseCase struct {
	repo repositories.UserRepository
}

func NewGetUserByIdUseCase(repo repositories.UserRepository) GetUserByIdUseCase {
	return &getUserByIdUseCase{repo: repo}
}

func (u *getUserByIdUseCase) Execute(ctx context.Context, input GetUserByIdInput) (*GetUserByIdOutput, error) {
	user, err := u.repo.FindById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	// Mapping User to GetUserByIdOutput
	// Hide sensitive data (e.g. password)
	result := &GetUserByIdOutput{
		ID:        user.ID,
		UserName:  user.UserName,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	// Other way to mapping struct using utils.MappingInterface
	//var result GetUserByIdOutput
	//err := utils.MappingInterface(user, &result)
	//if err != nil {
	//	return nil, err
	//}

	return result, nil
}
