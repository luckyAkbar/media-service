-- +migrate Up notransaction
ALTER TABLE  videos ADD COLUMN IF NOT EXISTS mime_type TEXT NOT NULL;

-- +migrate Down
ALTER TABLE videos DROP COLUMN IF EXISTS mime_type;