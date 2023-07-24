package repository

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type ArithmeticRepository struct {
	db *sql.DB
}

func NewArithmeticRepository(db *sql.DB) *ArithmeticRepository {
	return &ArithmeticRepository{db}
}

func (repo *ArithmeticRepository) AddToHistory(a, b, answer int32, operation string) error {
	//argsにValues()の引数に指定した値の配列が返ってくる
	queryString, args, err := sq.Insert("arith_history").Columns("date", "a", "b", "answer", "operation").Values(time.Now(), a, b, answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
