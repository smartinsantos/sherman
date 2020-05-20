-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE security_tokens (
   id               char(36)        NOT NULL,
   user_id          char(36)        NOT NULL,
   token            char(255)       NOT NULL,
   type             varchar(12)     NOT NULL,
   created_at       datetime        NOT NULL,
   updated_at       datetime        NOT NULL,
   PRIMARY KEY(id),
   FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE security_tokens ADD UNIQUE INDEX (user_id, type);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE security_tokens;
