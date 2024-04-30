CREATE TABLE IF NOT EXISTS tokens (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    status boolean NOT NULL DEFAULT true,

    name varchar(20) NOT NULL,
    remark varchar(200) NOT NULL DEFAULT '',
    token varchar(100) NOT NULL,

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);