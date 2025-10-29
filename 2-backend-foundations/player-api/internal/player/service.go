package player

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type RegisterParams struct {
	Username string
	Email    string
	Password string
}

type Service struct {
	players PlayerRepo
}

func NewService(pr PlayerRepo) *Service {
	return &Service{
		players: pr,
	}
}

func (s *Service) CreatePlayer(ctx context.Context, rp RegisterParams) (string, error) {
	if rp.Username == "" || rp.Email == "" || len(rp.Password) == 0 {
		return "", errors.New("username, email, and password required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(rp.Password), 12)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	p := Player{
		Username:       rp.Username,
		Email:          rp.Email,
		HashedPassword: hash,
	}

	return s.players.Create(ctx, p)
}

func (s *Service) GetPlayer(ctx context.Context, id string) (Player, error) {
	if id == "" {
		return Player{}, errors.New("id required")
	}
	return s.players.Get(ctx, id)
}
