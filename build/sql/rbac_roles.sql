CREATE TABLE IF NOT EXISTS rbac_roles (
    id bigserial PRIMARY KEY,
    status boolean NOT NULL DEFAULT true,

    name varchar(20) NOT NULL,
    description varchar(255) NOT NULL DEFAULT '',

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);