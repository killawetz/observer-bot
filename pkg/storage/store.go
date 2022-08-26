package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Repository struct {
	DBConn *pgx.Conn // for v5??
	DBPool *pgxpool.Pool
}

func New(connString string) *Repository {
	fmt.Println(connString)
	con, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Panic(err)
	}
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Panic(err)
	}
	return &Repository{
		DBConn: con,
		DBPool: pool,
	}
}

func (repo Repository) Conn() *pgx.Conn {
	return repo.DBConn
}
