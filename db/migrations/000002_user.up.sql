CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_id INT,
    email VARCHAR(255),
    password VARCHAR(255),

    CONSTRAINT USER_EMAIL_UNIQ UNIQUE (email),
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES persons(id)
);