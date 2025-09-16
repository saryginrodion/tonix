-- 7 - file_alter_timestamp up query
ALTER TABLE file
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
