package repository

import (
	"L0/common"
	"L0/config"
	"L0/database"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Repository struct {
	pool  *pgxpool.Pool
	mu    sync.RWMutex
	cache map[int][]byte
}

func New(cfg config.Config) *Repository {
	pool, err := database.ConnectDB(cfg)
	if err != nil {
		fmt.Println("Cann't connect to DB - ", err)
	}
	cache := make(map[int][]byte)
	return &Repository{pool: pool, cache: cache}
}

func (r *Repository) Get(id int) ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	bytes, ok := r.cache[id]
	if !ok {
		fmt.Println("ye")
		err := r.pool.QueryRow(context.Background(), "SELECT model FROM users WHERE id=$1", id).Scan(&bytes)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.ErrNotFound
		}
		if err != nil {
			return nil, err
		}
	}
	return bytes, nil
}
