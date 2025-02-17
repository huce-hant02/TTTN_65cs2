package edit_history

import (
	"context"
	"mono-base/internal/repositories"
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
	ModifierId *uint      `json:"modifierId"`
	Note       *string    `json:"note"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type GetInfoInputUseCase interface {
	Execute(ctx context.Context, input GetInfoInput) ([]GetInfoOutput, error)
}

type getInfoInputUseCase struct {
	repo repositories.EditHistoryRepository
}

func (g getInfoInputUseCase) Execute(ctx context.Context, input GetInfoInput) ([]GetInfoOutput, error) {
	history, err := g.repo.FindByModelId(ctx, input.ModelId)
	if err != nil {
		return nil, err
	}
	if history == nil {
		return []GetInfoOutput{}, nil
	}
	var results []GetInfoOutput
	for _, his := range history {
		result := GetInfoOutput{
			ID:         his.ID,
			ModelType:  his.ModelType,
			ModelId:    his.ModelId,
			Data:       his.Data,
			Active:     his.Active,
			ModifierId: his.ModifierId,
			Note:       his.Note,
			CreatedAt:  his.CreatedAt,
			UpdatedAt:  his.UpdatedAt,
		}
		results = append(results, result)
	}
	return results, nil
}

func NewGetInfoInputUseCase(repo repositories.EditHistoryRepository) GetInfoInputUseCase {
	return &getInfoInputUseCase{repo: repo}
}
