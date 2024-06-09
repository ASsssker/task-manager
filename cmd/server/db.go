package server

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func getDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.New("database connect timeout")
	}

	return db, nil
}