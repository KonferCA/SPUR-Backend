-- +goose Up
-- +goose StatementBegin
CREATE TABLE startups (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX idx_startups_owner_name ON startups (owner_id, name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE startups;
-- +goose StatementEnd