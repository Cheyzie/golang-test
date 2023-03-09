CREATE TABLE users_refresh_tokens
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    token VARCHAR(255) NOT NULL,
    session_name VARCHAR(255) NOT NULL
);