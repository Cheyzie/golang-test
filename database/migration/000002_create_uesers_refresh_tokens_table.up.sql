CREATE TABLE users_refresh_tokens
(
    id BIGINT PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    token VARCHAR(255) NOT NULL,
    details VARCHAR(255) NOT NULL
);