-- +goose Up
CREATE INDEX users_created_at_idx ON posts (created_at);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP INDEX users_created_at_idx;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
