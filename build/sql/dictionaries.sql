-- only store the low frequency updated or no need to update data
CREATE TABLE dictionaries (
    id bigserial PRIMARY KEY,

    category varchar(20),
    key varchar(20),
    value varchar(100),

    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
CREATE INDEX idx_category ON dictionaries USING btree (category);