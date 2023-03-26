-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category
(
    id         BIGINT      NOT NULL,
    nama       VARCHAR(256),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       image_url VARCHAR(255),
                       release_year INTEGER,
                       price VARCHAR(10),
                       total_page INTEGER,
                       thickness VARCHAR(10),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       category_id INTEGER
);

-- +migrate StatementEnd