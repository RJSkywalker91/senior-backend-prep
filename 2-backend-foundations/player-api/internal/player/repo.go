package player

import (
	"context"
)

type PlayerRepo interface {
	Create(ctx context.Context, p Player) (string, error)
	Get(ctx context.Context, id string) (Player, error)
}
