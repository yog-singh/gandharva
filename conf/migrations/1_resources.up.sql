CREATE TABLE IF NOT EXISTS resources (
    id                      UUID DEFAULT uuid_generate_v4(),
    name                    VARCHAR(255) NOT NULL,
    url                     VARCHAR(512) NOT NULL,
    request_method          VARCHAR(32) NOT NULL,
    ping_interval_in_mins   INTEGER NOT NULL,
    status                  VARCHAR(64) NOT NULL,
    last_checked_at         TIMESTAMP,
    created_at              TIMESTAMP,
    CONSTRAINT  pk_resources PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS resource_heartbeats (
    id                      UUID DEFAULT uuid_generate_v4(),
    resource_id             UUID NOT NULL,
    status_code             INTEGER NOT NULL,
    response_body           TEXT,
    latency                 BIGINT NOT NULL,
    created_at              TIMESTAMP,
    CONSTRAINT  pk_resource_heartbeats PRIMARY KEY (id),
    CONSTRAINT  fk_resource_heartbeats_resource FOREIGN KEY (resource_id) REFERENCES resources (id)
);
