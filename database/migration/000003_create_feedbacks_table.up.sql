CREATE TABLE feedbacks
(
    id SERIAL PRIMARY KEY NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    feedback_text VARCHAR(255) NOT NULL,
    source VARCHAR(255) NOT NULL
)