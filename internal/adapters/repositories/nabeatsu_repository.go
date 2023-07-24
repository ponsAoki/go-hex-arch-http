package repository

import (
	"database/sql"
	entity "hex/internal/adapters/entities/nabeatsu"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type NabeatsuRepository struct {
	db *sql.DB
}

func NewNabeatsuRepository(db *sql.DB) *NabeatsuRepository {
	return &NabeatsuRepository{db}
}

func (repo *NabeatsuRepository) CreateNabeatsu(nabeatsu *entity.Nabeatsu) error {
	queryString, args, err := sq.Insert("nabeatsu").Columns("date", "a", "b", "answer", "operation").Values(time.Now(), nabeatsu.A, nabeatsu.B, nabeatsu.Answer, nabeatsu.Operation).ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
