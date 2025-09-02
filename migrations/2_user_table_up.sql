-- 2 - user_table up query
CREATE EXTENSION IF NOT EXISTS pgcrypto;
-- migren:split
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(320) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    username VARCHAR(256) NOT NULL UNIQUE,
    displayable_name VARCHAR(256) NOT NULL DEFAULT '',
    description VARCHAR(256) NOT NULL DEFAULT '',
    email_verified BOOLEAN NOT NULL DEFAULT false,
    balance INTEGER NOT NULL DEFAULT 0,
    withdrawal_balance INTEGER NOT NULL DEFAULT 0,
    last_withdrawal_card VARCHAR(255) NOT NULL DEFAULT ''
);
