-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE security_tokens (
   id               char(36)        NOT NULL,
   user_id          char(36)        NOT NULL,
   token            char(64)        NOT NULL,
   created_at       datetime        NOT NULL,
   updated_at       datetime        NOT NULL,
   PRIMARY KEY(id),
   FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE security_tokens;
