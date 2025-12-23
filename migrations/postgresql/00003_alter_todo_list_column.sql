-- +goose Up
-- +goose StatementBegin
DELETE FROM todo_list WHERE user_id IS NULL;
ALTER TABLE todo_list
DROP CONSTRAINT IF EXISTS todo_list_user_id_fkey;

ALTER TABLE todo_list
ALTER COLUMN user_id SET NOT NULL;

ALTER TABLE todo_list
ADD CONSTRAINT todo_list_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todo_list
DROP CONSTRAINT IF EXISTS todo_list_user_id_fkey;

ALTER TABLE todo_list
ALTER COLUMN user_id DROP NOT NULL;

ALTER TABLE todo_list
ADD CONSTRAINT todo_list_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES users(id);
-- +goose StatementEnd
