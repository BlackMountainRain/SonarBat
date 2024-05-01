CREATE TYPE ip_with_info AS (
    ip INET,
    is_private BOOLEAN,
    is_ipv6 BOOLEAN
);

CREATE TABLE IF NOT EXISTS hosts (
   id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
   status boolean DEFAULT true NOT NULL,

   name varchar(250) NOT NULL,
   live_at timestamptz NULL,
   is_agent_installed boolean DEFAULT false NOT NULL,
   agent_version varchar(10) NULL,

   ips ip_with_info[] NOT NULL,
   net_type smallint NULL,

   additions jsonb NULL,

   created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
   updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);
CREATE INDEX idx_hosts_name ON hosts USING btree (name);
COMMENT ON COLUMN hosts.name IS 'endpoint name, maybe hostname or ip';
COMMENT ON COLUMN hosts.is_agent_installed IS 'this host is installed agent or not, only installed agent can run tasks';
COMMENT ON COLUMN hosts.agent_version IS 'agent version if installed';
COMMENT ON COLUMN hosts.net_type IS 'network type of this host: 2: dual-stack, 4: IPv4 only, 6: IPv6 only';
COMMENT ON COLUMN hosts.live_at IS 'last live time of this host';