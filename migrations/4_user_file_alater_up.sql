-- 4 - user_file_alater up query
ALTER TABLE users
ADD COLUMN avatar_id UUID REFERENCES file(id) NULL;
-- migren:split
ALTER TABLE users
ADD COLUMN identity_photo_id UUID REFERENCES file(id) NULL;
