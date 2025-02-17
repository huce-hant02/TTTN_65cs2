package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	constants2 "mono-base/pkg/constants"
	error2 "mono-base/pkg/error"
	"net/http"
)

func GetContextTransaction(ctx context.Context) *sql.Tx {
	if ctx.Value(constants2.ContextKeyDBTransaction) != nil {
		return ctx.Value(constants2.ContextKeyDBTransaction).(*sql.Tx)
	}
	return nil
}

func txSelect(tx *sql.Tx, dest interface{}, query string, args ...interface{}) error {
	rows, err := tx.Query(query, args...)
	if err != nil {
		return err
	}
	err = sqlx.StructScan(rows, dest)
	if err != nil {
		return err
	}
	return nil
}

func Select(ctx context.Context, db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	tx := GetContextTransaction(ctx)
	if tx != nil {
		err := txSelect(tx, dest, query, args...)
		if err != nil {
			return err
		}
	} else {
		err := db.Select(dest, query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}

func Insert(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) (*int, error) {
	tx := GetContextTransaction(ctx)
	queryS := fmt.Sprintf("%s %s", query, "RETURNING id")
	var id int
	var err error
	if tx != nil {
		err = tx.QueryRow(queryS, args...).Scan(&id)
	} else {
		err = db.QueryRow(queryS, args...).Scan(&id)
	}
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			return nil, error2.NewError(http.StatusBadRequest, mErr.Error(), "data_invalid")
		}
		return nil, err
	}
	return &id, nil
}

func InsertMultiple(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) ([]int, error) {
	tx := GetContextTransaction(ctx)
	queryS := fmt.Sprintf("%s %s", query, "RETURNING id")
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.Query(queryS, args...)
	} else {
		rows, err = db.Query(queryS, args...)
	}
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			return nil, error2.NewError(http.StatusBadRequest, mErr.Error(), "data_invalid")
		}
		return nil, err
	}
	ids := make([]int, 0)
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func Update(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) error {
	tx := GetContextTransaction(ctx)
	var err error
	if tx != nil {
		_, err = tx.Exec(query, args...)
	} else {
		_, err = db.Exec(query, args...)
	}
	return err
}

func Delete(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) error {
	tx := GetContextTransaction(ctx)
	var err error
	if tx != nil {
		_, err = tx.Exec(query, args...)
	} else {
		_, err = db.Exec(query, args...)
	}
	return err
}
