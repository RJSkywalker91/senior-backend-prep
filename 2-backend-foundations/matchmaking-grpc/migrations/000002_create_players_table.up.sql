-- Players table: bcrypt hash stored as bytea; unique username/email
CREATE TABLE IF NOT EXISTS players (
    id           BIGSERIAL PRIMARY KEY,
    username     TEXT NOT NULL,
    email        CITEXT NOT NULL,
    hash         BYTEA NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Enforce uniqueness
CREATE UNIQUE INDEX IF NOT EXISTS ux_players_username ON players (username);
CREATE UNIQUE INDEX IF NOT EXISTS ux_players_email ON players (email);

-- Fast lookup by username (used on login)
CREATE INDEX IF NOT EXISTS ix_players_username ON players (username);