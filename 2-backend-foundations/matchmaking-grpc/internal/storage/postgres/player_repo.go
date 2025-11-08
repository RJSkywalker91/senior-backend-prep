package postgres

import (
	"context"
	"errors"
	"matchmaking/internal/player"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNoRows = errors.New("no player found")
var ErrConflict = errors.New("database conflict occurred")

type PlayerRepoPG struct {
	db *pgxpool.Pool
}

func NewPlayerRepoPG(db *pgxpool.Pool) *PlayerRepoPG {
	return &PlayerRepoPG{db: db}
}

func (r *PlayerRepoPG) Create(ctx context.Context, p player.Player) (int64, error) {
	var id int64
	query := `INSERT INTO players(username, email, hash) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query,
		p.Username, p.Email, string(p.HashedPassword),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PlayerRepoPG) Get(ctx context.Context, username string) (*player.Player, error) {
	var p player.Player
	query := `select id, username, email, hash, created_at from players where username=$1`
	err := r.db.QueryRow(ctx, query, username).Scan(&p.Id, &p.Username, &p.Email, &p.HashedPassword, &p.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRows
		}
		return nil, err
	}
	return &p, nil
}
