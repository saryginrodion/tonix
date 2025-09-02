-- 4 - user_file_alater down query
ALTER TABLE users
DROP COLUMN avatar_id;
-- migren:split
ALTER TABLE users
DROP COLUMN identity_photo_id;
