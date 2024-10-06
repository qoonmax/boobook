-- +goose Up
CREATE INDEX users_last_name_first_name_idx ON users (last_name text_pattern_ops, first_name text_pattern_ops);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP INDEX users_last_name_first_name_idx;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
