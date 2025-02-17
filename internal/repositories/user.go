package repositories

import (
	"context"
	"mono-base/internal/entities"
)

type UserRepository interface {
	FindById(ctx context.Context, id int) (*entities.User, error)
}
