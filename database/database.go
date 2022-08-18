package database

import (
	"context"

	"europe/config"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	*sqlx.DB
	schemaName string
}

func NewClient(cfg *config.Config) (*Client, error) {
	db, err := sqlx.Open("pgx", cfg.DB.URL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	return &Client{
		db,
		cfg.DB.SchemaName,
	}, nil
}

const statusCheckQuery = `SELECT true`

func (db *Client) StatusCheck(ctx context.Context) error {
	var tmp bool

	return db.QueryRowContext(ctx, statusCheckQuery).Scan(&tmp)
}
