-- +goose Up
CREATE TYPE gender AS ENUM ('male', 'female', 'other');

CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password      TEXT         NOT NULL,
    first_name    VARCHAR(255) NOT NULL,
    last_name     VARCHAR(255) NOT NULL,
    date_of_birth DATE         NOT NULL,
    gender        gender       NOT NULL,
    interests     TEXT         NOT NULL,
    city          VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;
DROP TYPE gender;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
