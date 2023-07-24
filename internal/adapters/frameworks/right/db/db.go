package db

import (
	"database/sql"
	"log"
)

type Adapter struct {
	DB *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}

	//test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failed: %v", err)
	}

	return &Adapter{DB: db}, nil
}

func (da *Adapter) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("db close failed: %v", err)
	}
}
