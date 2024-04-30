CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    status boolean NOT NULL DEFAULT true,

    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,

    roles BIGINT[],

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
CREATE INDEX idx_username ON users USING btree (username);
CREATE UNIQUE INDEX idx_email ON users USING btree (email);