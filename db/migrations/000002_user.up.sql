CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    person_id INT,
    email VARCHAR(255),
    password VARCHAR(255),

    CONSTRAINT USER_EMAIL_UNIQ UNIQUE (email),
    CONSTRAINT person_id_fk FOREIGN KEY (person_id) REFERENCES persons(id)
);