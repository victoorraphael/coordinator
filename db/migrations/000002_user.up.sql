CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) references persons (email),
    password VARCHAR(255),

    CONSTRAINT USER_EMAIL_UNIQ UNIQUE (email)
);