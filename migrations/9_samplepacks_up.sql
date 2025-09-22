-- 9 - samplepacks up query
CREATE TYPE status_type AS ENUM ('draft', 'moderation', 'published');

-- migren:split
CREATE TABLE samplepacks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    cost INTEGER  NOT NULL,
    preview_id UUID REFERENCES files(id),
    icon_id UUID REFERENCES files(id),
    author_id UUID REFERENCES users(id) NOT NULL,
    purchase_count INTEGER NOT NULL DEFAULT 0,
    status status_type NOT NULL DEFAULT 'draft',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    likes INTEGER NOT NULL DEFAULT 0
);

-- migren:split
CREATE TABLE samplepack_tags (
    name TEXT NOT NULL,
    samplepack_id UUID NOT NULL,
    PRIMARY KEY (name, samplepack_id)
);

-- migren:split
CREATE TABLE purchased_samplepacks (
    user_id UUID REFERENCES users(id) NOT NULL,
    samplepack_id UUID REFERENCES samplepacks(id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, samplepack_id)
)

-- migren:split
CREATE TABLE liked_samplepacks (
    user_id UUID REFERENCES users(id) NOT NULL,
    samplepack_id UUID REFERENCES samplepacks(id) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, samplepack_id)
)
