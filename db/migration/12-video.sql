-- +migrate Up notransaction
CREATE TABLE IF NOT EXISTS videos (
    id BIGINT PRIMARY KEY,
    filename TEXT NOT NULL,
    title TEXT DEFAULT '',
    extention VARCHAR(100) NOT NULL,
    size_bytes BIGINT NOT NULL,
    length_seconds BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS videos;