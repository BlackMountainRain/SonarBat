CREATE TABLE IF NOT EXISTS rbac_objects (
    id bigserial PRIMARY KEY,
    status boolean NOT NULL DEFAULT true,

    value varchar(50) NOT NULL,

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);