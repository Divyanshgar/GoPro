-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL References users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL References feeds(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;