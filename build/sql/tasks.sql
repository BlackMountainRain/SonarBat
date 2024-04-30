CREATE TABLE IF NOT EXISTS tasks (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    status boolean DEFAULT true NOT NULL,

    name varchar(100) NOT NULL,
    comment varchar(255) NOT NULL DEFAULT '',

    updated_by uuid NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
CREATE INDEX idx_name ON tasks USING btree (name);
COMMENT ON COLUMN tasks.updated_by IS 'latest updated user id';
COMMENT ON COLUMN tasks.created_by IS 'creator user id';