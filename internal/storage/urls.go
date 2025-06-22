package storage

import (
	"context"
)

func (s *Storage) Save(ctx context.Context, short, original string) error {
	_, err := s.DB.Exec(ctx, `INSERT INTO urls (short, original) VALUES ($1, $2)`, short, original)
	return err
}

func (s *Storage) Find(ctx context.Context, short string) (string, error) {
	var original string
	err := s.DB.QueryRow(ctx, `SELECT original FROM urls WHERE short = $1`, short).Scan(&original)
	return original, err
}
