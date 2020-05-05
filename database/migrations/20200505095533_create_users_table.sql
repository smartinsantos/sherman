-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
   id               char(36)        NOT NULL,
   first_name       varchar(100)    NOT NULL,
   last_name        varchar(100)    NOT NULL,
   email_address    varchar(100)    NOT NULL UNIQUE,
   password         varchar(100)    NOT NULL,
   active           bit             NOT NULL,
   created_at       datetime        NOT NULL,
   updated_at       datetime        NOT NULL,
   PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
