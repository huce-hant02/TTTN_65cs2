package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type BaseRepository[T any] interface {
	GetByCondition(ctx context.Context, condition *CommonCondition) (*Pagination[T], error)
	GetById(ctx context.Context, id string) (*T, error)
	GetByIds(ctx context.Context, ids []string) ([]*T, error)
	Create(ctx context.Context, entity *T) (*T, error)
	Upsert(ctx context.Context, entity *T, uniqueColumns []string, updateColumns []string) (*T, error)
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, ids []string) error
	GetDBName() string
	GetTable() string
	GetDBContext() *sqlx.DB
}
