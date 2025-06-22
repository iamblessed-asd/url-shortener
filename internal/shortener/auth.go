// internal/shortener/auth.go
package shortener

import (
	"context"
	"log"
	"url-shortener/internal/auth"
)

func (s *Service) RegisterUser(ctx context.Context, username, password string) error {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}
	log.Println("Added user: ", username)
	return s.Repo.CreateUser(ctx, username, hash)
}

func (s *Service) AuthenticateUser(ctx context.Context, username, password string) (int, error) {
	id, hash, err := s.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, err
	}
	if id == 0 || !auth.CheckPasswordHash(password, hash) {
		return 0, ErrInvalidCredentials
	}
	return id, nil
}
