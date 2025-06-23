package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func (s *Storage) CreateUser(ctx context.Context, username, passwordHash string) error {
	_, err := s.DB.Exec(ctx, "INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, passwordHash)
	log.Println("Inserted user ", username, " into DB")
	return err
}

func (s *Storage) GetUserByUsername(ctx context.Context, username string) (int, string, error) {
	var id int
	var hash string
	err := s.DB.QueryRow(ctx, "SELECT id, password_hash FROM users WHERE username=$1", username).Scan(&id, &hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", nil
		}
		return 0, "", err
	}
	return id, hash, nil
}
