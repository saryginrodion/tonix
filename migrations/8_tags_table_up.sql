-- 8 - tags_table up query
CREATE TYPE tag_type AS ENUM ('unsorted', 'instrument', 'genre');

-- migren:split
CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name CHAR(255) UNIQUE NOT NULL,
    type tag_type NOT NULL,
    usages INTEGER NOT NULL DEFAULT 1
);
