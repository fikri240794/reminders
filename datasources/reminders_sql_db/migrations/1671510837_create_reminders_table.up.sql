CREATE TABLE reminders (
    id CHAR(36) NOT NULL,
    note TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    PRIMARY KEY (id)
)