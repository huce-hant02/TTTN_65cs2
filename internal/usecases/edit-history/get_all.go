package edit_history

import (
	"context"
	"go.uber.org/zap"
	"mono-base/internal/repositories"
	"mono-base/pkg/logger"
	"time"
)

type GetInfoInput struct {
	ModelId uint `json:"modelId"`
}

type GetInfoOutput struct {
	ID         uint       `json:"id"`
	ModelType  string     `json:"modelType"`
	ModelId    uint       `json:"modelId"`
	Data       string     `json:"data"`
	Active     *bool      `json:"active"`
	ModifierId *int       `json:"modifierId"`
	Note       *string    `json:"note"`
	Author     *string    `json:"author"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type GetInfoInputUseCase interface {
	Execute(ctx context.Context, input GetInfoInput) ([]GetInfoOutput, error)
}

type getInfoInputUseCase struct {
	repo repositories.EditHistoryRepository
	user repositories.UserRepository
}

func NewGetInfoInputUseCase(
	repo repositories.EditHistoryRepository,
	user repositories.UserRepository,
) GetInfoInputUseCase {
	return &getInfoInputUseCase{
		repo: repo,
		user: user,
	}
}

func (g getInfoInputUseCase) Execute(ctx context.Context, input GetInfoInput) ([]GetInfoOutput, error) {
	ctxLogger := logger.NewLogger(ctx)
	history, err := g.repo.FindByModelId(ctx, input.ModelId)
	if err != nil {
		ctxLogger.Error("find history by modelId", zap.Error(err))
		return nil, err
	}
	if history == nil {
		ctxLogger.Info("history is nil", zap.Any("history", history))
		return []GetInfoOutput{}, nil
	}
	userId := history[0].ModifierId
	author, err := g.user.FindById(ctx, *userId)
	if err != nil {
		ctxLogger.Error("find user by id", zap.Error(err))
		return nil, err
	}
	if author == nil {
		ctxLogger.Info("author is nil", zap.Any("author", author))
	}

	var results []GetInfoOutput
	for _, his := range history {
		result := GetInfoOutput{
			ID:         his.ID,
			ModelType:  his.ModelType,
			ModelId:    uint(his.ModelId),
			Data:       his.Data,
			Active:     his.Active,
			ModifierId: his.ModifierId,
			Author:     &author.UserName,
			Note:       his.Note,
			CreatedAt:  his.CreatedAt,
			UpdatedAt:  his.UpdatedAt,
		}
		results = append(results, result)
	}
	return results, nil
}
