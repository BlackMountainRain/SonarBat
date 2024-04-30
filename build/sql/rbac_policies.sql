CREATE TABLE IF NOT EXISTS rbac_policies (
    id bigserial PRIMARY KEY,

    role varchar(20) NULL,
    obj varchar(20) NULL,
    act varchar(20) NULL,
    uri varchar(20) NULL,

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON COLUMN rbac_policies.role IS 'role or user';
COMMENT ON COLUMN rbac_policies.obj IS 'module or api';
COMMENT ON COLUMN rbac_policies.act IS 'action or method';
COMMENT ON COLUMN rbac_policies.uri IS 'uri or path';