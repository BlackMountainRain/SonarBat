CREATE TABLE IF NOT EXISTS host_blacklist (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    host_id uuid NOT NULL,

    reason varchar(50) NOT NULL,

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
CREATE INDEX idx_host_blacklist_host_id ON host_blacklist USING btree (host_id);