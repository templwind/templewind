-- +goose Up
-- +goose StatementBegin
PRAGMA foreign_keys = ON;

-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts(
    id text PRIMARY KEY DEFAULT ('r' || lower(hex(randomblob(7)))) NOT NULL,
    company_name text,
    address_1 text,
    address_2 text,
    city text,
    state_province text,
    postal_code text,
    country text,
    phone text,
    email text,
    website text,
    primary_user_id text NOT NULL,
    created_at text DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at text DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_types(
    id text PRIMARY KEY NOT NULL,
    type_name text NOT NULL,
    description text NOT NULL
);

-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id text PRIMARY KEY DEFAULT ('r' || lower(hex(randomblob(7)))) NOT NULL,
    name text DEFAULT '' NOT NULL,
    username text NOT NULL,
    email text DEFAULT '' NOT NULL UNIQUE,
    email_visibility boolean DEFAULT FALSE NOT NULL,
    last_reset_sent_at text DEFAULT '' NOT NULL,
    last_verification_sent_at text DEFAULT '' NOT NULL,
    password_hash text NOT NULL,
    token_key text NOT NULL,
    verified boolean DEFAULT FALSE NOT NULL,
    avatar text DEFAULT '' NOT NULL,
    type_id text DEFAULT 'USER_TYPE_ADMIN' NOT NULL,
    created_at text DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at text DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (type_id) REFERENCES user_types(id) ON DELETE SET NULL
);

-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_accounts(
    user_id text NOT NULL,
    account_id text NOT NULL,
    PRIMARY KEY (user_id, account_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
);

CREATE INDEX idx_user_accounts_user_id ON user_accounts(user_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_accounts;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS user_types;

DROP TABLE IF EXISTS accounts;

-- +goose StatementEnd
