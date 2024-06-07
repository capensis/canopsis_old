BEGIN;

CREATE TABLE IF NOT EXISTS action_log(
    id            BIGSERIAL PRIMARY KEY,
    type          SMALLINT,
    value_type    VARCHAR(255) NOT NULL,
    value_id      VARCHAR(255) NOT NULL,
    author        VARCHAR(255),
    time          TIMESTAMP NOT NULL DEFAULT NOW(),
    data          JSONB
);

CREATE INDEX IF NOT EXISTS action_log_value_type_idx ON action_log (value_type);
CREATE INDEX IF NOT EXISTS action_log_value_id_idx ON action_log (value_id);
CREATE INDEX IF NOT EXISTS action_log_value_author_idx ON action_log (author);
CREATE INDEX IF NOT EXISTS action_log_time_idx ON action_log (time);

COMMIT;
