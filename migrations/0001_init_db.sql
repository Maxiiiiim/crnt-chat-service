-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id          uuid PRIMARY KEY,
    meta        json,
    last_active TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS dialogs
(
    id              uuid PRIMARY KEY,
    meta            json,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_message_id uuid                     NOT NULL,
    personal        bool                     NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS messages
(
    id                 uuid PRIMARY KEY,
    dialog_id          uuid                     NOT NULL,
    sender_id          uuid                     NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    sent_at            TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    content_type       INT                      NOT NULL DEFAULT 0,
    content_text       TEXT                     NOT NULL DEFAULT '',
    content_additional bytea[]                  NOT NULL DEFAULT ARRAY []::bytea[],
    content_meta       json                     NOT NULL DEFAULT '{}',
    is_deleted         bool                     NOT NULL DEFAULT FALSE,
    reply_to_id        uuid,
    replies_count      INT                      NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS dialog_members
(
    dialog_id            uuid NOT NULL,
    user_id              uuid NOT NULL,
    is_owner             bool NOT NULL DEFAULT FALSE,
    last_read_message_id uuid NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    PRIMARY KEY (dialog_id, user_id)
);
-- +goose StatementEnd