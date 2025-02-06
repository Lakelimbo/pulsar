-- +goose Up
CREATE TABLE dummies (
    id UUID DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    --
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE dummies;