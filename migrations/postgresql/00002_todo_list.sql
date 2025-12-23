-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todo_list(
    id uuid primary key,
    user_id uuid references users(id),
    title varchar(255) not null,
    completed boolean default FALSE,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_list;
-- +goose StatementEnd
