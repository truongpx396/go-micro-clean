-- CREATE TABLE IF NOT EXISTS test (id SERIAL PRIMARY KEY, name VARCHAR(50));

BEGIN;

-- structure setup

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- data setup

INSERT INTO users (username, email) 
VALUES ('user1', 'user1@example.com');

INSERT INTO users (username, email) 
VALUES ('user2', 'user2@example.com');

COMMIT;