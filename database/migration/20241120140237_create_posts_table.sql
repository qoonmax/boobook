-- +goose Up

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE posts ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE posts;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
