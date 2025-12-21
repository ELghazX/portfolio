-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Up
CREATE TYPE project_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE exp_type AS ENUM ('work', 'education', 'certification');

CREATE TABLE projects (
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    slug        TEXT UNIQUE NOT NULL,
    summary     TEXT NOT NULL,
    content     TEXT NOT NULL,
    repo_url    TEXT,
    live_url    TEXT,
    tags        TEXT[], 
    image_url   TEXT,
    status      project_status DEFAULT 'draft',
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE experiences (
    id          SERIAL PRIMARY KEY,
    type        exp_type NOT NULL,
    title       TEXT NOT NULL,
    institution TEXT NOT NULL,
    location    TEXT,
    start_date  DATE NOT NULL,
    end_date    DATE,
    description TEXT,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    slug        TEXT UNIQUE NOT NULL,
    summary     TEXT NOT NULL,
    content     TEXT NOT NULL,
    tags        TEXT[],
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
drop TYPE IF EXISTS project_status;
drop TYPE IF EXISTS exp_type;
drop TABLE IF EXISTS projects;
drop TABLE IF EXISTS experiences;
drop TABLE IF EXISTS posts;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
