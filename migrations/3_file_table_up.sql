-- 3 - file_table up query
CREATE TABLE IF NOT EXISTS file (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    filename VARCHAR(320),
    mimetype VARCHAR(320),
    author_id UUID REFERENCES users(id)
);
