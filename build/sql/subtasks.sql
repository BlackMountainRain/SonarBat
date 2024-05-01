CREATE TABLE IF NOT EXISTS subtasks (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    task_id uuid NOT NULL,
    status boolean DEFAULT true NOT NULL,

    detection_type smallint NOT NULL,
    params jsonb NULL,

    src_ep_filter_strategy jsonb NULL,
    src_ep_sel_strategy jsonb NULL,
    src_ep_sel_num int NOT NULL,

    dst_ep_filter_strategy jsonb NULL,
    dst_ep_sel_strategy jsonb NULL,
    dst_ep_sel_num int NOT NULL,

    updated_by uuid NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
CREATE INDEX idx_subtasks_task_id ON subtasks USING btree (task_id);
CREATE INDEX idx_subtasks_detection_type ON subtasks USING btree (detection_type);
COMMENT ON COLUMN subtasks.detection_type IS 'detect type: 1: Ping, 2: TcpPing, 3: UdpPing, 4: Http, 5: Tcp, 6: Udp, 7: Dns, 8: Mtr';
COMMENT ON COLUMN subtasks.src_ep_filter_strategy IS '{key: value} to filter endpoints';
COMMENT ON COLUMN subtasks.src_ep_sel_strategy IS '{key: value} to select endpoints from filtered pool';
COMMENT ON COLUMN subtasks.updated_by IS 'latest updated user id';
COMMENT ON COLUMN subtasks.created_by IS 'creator user id';