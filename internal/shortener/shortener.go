package shortener

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"strings"
)

type Repository interface {
	Save(ctx context.Context, short, original string) error
	Find(ctx context.Context, short string) (string, error)
	CreateUser(ctx context.Context, username, hash string) error
	GetUserByUsername(ctx context.Context, username string) (int, string, error)
}

type Service struct {
	Repo Repository
}

func New(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) Shorten(ctx context.Context, original string) (string, error) {
	short := generateCode()
	err := s.Repo.Save(ctx, short, original)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (s *Service) Resolve(ctx context.Context, short string) (string, error) {
	return s.Repo.Find(ctx, short)
}

func generateCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}
