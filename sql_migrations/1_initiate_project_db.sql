-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE IF NOT EXISTS auth_user_pkey_seq;
CREATE TABLE "auth_user"
(
    id             BIGINT NOT NULL             DEFAULT nextval('auth_user_pkey_seq'::regclass),
    full_name      VARCHAR(50) NOT NULL,
    username       VARCHAR(10) NOT NULL,
    password       VARCHAR(256) NOT NULL,
    created_by     BIGINT,
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by     BIGINT,
    updated_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted        BOOLEAN DEFAULT FALSE,
    CONSTRAINT pk_auth_user_id PRIMARY KEY (id),
    CONSTRAINT uq_auth_user_username UNIQUE (username)
);

-- +migrate StatementEnd