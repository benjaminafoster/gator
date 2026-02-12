-- +goose up
ALTER TABLE feeds ADD COLUMN last_fetched TIMESTAMP;

-- +goose down
ALTER TABLE feeds DROP COLUMN last_fetched;