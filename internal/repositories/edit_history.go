package repositories

import (
	"context"
	"mono-base/internal/entities"
)

type EditHistoryRepository interface {
	FindByModelId(ctx context.Context, id uint) ([]entities.CdioEditHistory, error)
}
