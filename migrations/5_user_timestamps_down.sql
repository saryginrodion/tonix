-- 5 - user_timestamps down query
ALTER TABLE users
DROP COLUMN created_at,
DROP COLUMN updated_at;
