-- 8 - tags_table up query
CREATE TYPE tag_type AS ENUM ('unsorted', 'instrument', 'genre');

-- migren:split
CREATE TABLE tags (
    name TEXT PRIMARY KEY NOT NULL,
    type tag_type NOT NULL,
    usages INTEGER NOT NULL DEFAULT 1
);
