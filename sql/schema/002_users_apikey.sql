-- +goose Up

Alter TABLE users
Add COLUMN api_key Char(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
Alter TABLE users
Drop COLUMN api_key;