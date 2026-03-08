-- +goose Up
-- +goose StatementBegin
ALTER TABLE events
ADD IF NOT EXISTS type TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE events
DROP COLUMN IF EXISTS type;
-- +goose StatementEnd