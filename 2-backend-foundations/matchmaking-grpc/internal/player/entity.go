package player

import "time"

type Player struct {
	Id             int64     `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}
