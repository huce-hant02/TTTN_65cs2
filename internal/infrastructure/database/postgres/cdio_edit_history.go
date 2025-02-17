package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"mono-base/internal/entities"
	"mono-base/internal/repositories"
	"mono-base/pkg/logger"
)

type editHistoryRepository struct {
	db    *sqlx.DB
	table string
}

func (e editHistoryRepository) FindByModelId(ctx context.Context, id uint) ([]entities.CdioEditHistory, error) {
	ctxLogger := logger.NewLogger(ctx)
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	db := psql.Select("*").
		From(e.table).
		Where(sq.Eq{"model_id": id}).
		Where(sq.Eq{"deleted_at": nil})
	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, err: %v", err)
		return nil, err
	}
	var users []entities.CdioEditHistory
	err = Select(ctx, e.db, &users, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed while select history, err: %v", err)
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users, nil
}

func NewEditHistoryRepository(db *sqlx.DB) repositories.EditHistoryRepository {
	return &editHistoryRepository{
		db:    db,
		table: "cdio_edit_histories",
	}
}
