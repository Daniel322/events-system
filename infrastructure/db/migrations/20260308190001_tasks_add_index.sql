-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS tasks_date
ON tasks (date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS tasks_date on tasks;
-- +goose StatementEnd