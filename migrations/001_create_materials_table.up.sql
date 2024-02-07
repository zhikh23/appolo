CREATE TABLE IF NOT EXISTS materials (
    id          BIGSERIAL       PRIMARY KEY,
    name        VARCHAR(256)    NOT NULL    DEFAULT '',
    description VARCHAR(256)    NOT NULL    DEFAULT '',
    tags        VARCHAR(32)[]   NOT NULL    DEFAULT '{}',
    url         VARCHAR(256)    NOT NULL,

    created_at  TIMESTAMP       NOT NULL    DEFAULT NOW(),

    CONSTRAINT url_not_empty
        CHECK (url != '')
);
