package postgres

import (
	"context"
	"errors"
	"playerapi/internal/player"
	"strconv"

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

func (r *PlayerRepoPG) Create(ctx context.Context, p player.Player) (string, error) {
	var id int
	err := r.db.QueryRow(ctx,
		`insert into players(username, email, hash)
		 values ($1, $2, $3)
		 returning id`,
		p.Username, p.Email, string(p.HashedPassword)).
		Scan(&id)

	if err != nil {
		return "", err
	}

	idStr := strconv.Itoa(id)
	return idStr, nil
}

func (r *PlayerRepoPG) Get(ctx context.Context, id string) (player.Player, error) {
	var p player.Player
	err := r.db.QueryRow(ctx,
		`select id, username, email, hash, created_at from players where id=$1`, id).
		Scan(&p.Id, &p.Username, &p.Email, &p.HashedPassword, &p.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return player.Player{}, ErrNoRows
		}
		return player.Player{}, err
	}
	return p, nil
}
