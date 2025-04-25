-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS events(
  id UUID DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users (id),
  info TEXT NOT NULL,
  date TIMESTAMP NOT NULL,
  notify_levels JSONB NOT NULL,
  providers JSONB NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT id_events_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
