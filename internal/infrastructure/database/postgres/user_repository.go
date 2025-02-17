package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"mono-base/internal/entities"
	"mono-base/internal/repositories"
	"mono-base/pkg/logger"
)

type userRepository struct {
	db    *sqlx.DB
	table string
}

func NewUserRepository(db *sqlx.DB) repositories.UserRepository {
	return &userRepository{
		db:    db,
		table: "users",
	}
}

func (u *userRepository) FindById(ctx context.Context, id int) (*entities.User, error) {
	ctxLogger := logger.NewLogger(ctx)
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	db := psql.Select("*").
		From(u.table).
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"deleted_at": nil})
	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, err: %v", err)
		return nil, err
	}
	var users []entities.User
	err = Select(ctx, u.db, &users, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed while select users, err: %v", err)
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}
