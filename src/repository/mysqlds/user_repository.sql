-- name: create-user
INSERT INTO users (
    id,
    first_name,
    last_name,
    email_address,
    password,
    active,
    created_at,
    updated_at
)
VALUES(?, ?, ?, ?, ?, ?, ?, ?);

-- name: find-user-by-email
SELECT * FROM users WHERE (email_address) = ? LIMIT 1;
