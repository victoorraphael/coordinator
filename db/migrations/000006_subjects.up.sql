CREATE TABLE IF NOT EXISTS subjects (
    id SERIAL,
    nome VARCHAR(255),
    professor_id int,

    PRIMARY KEY (id),
    FOREIGN KEY (professor_id) REFERENCES professor (id)
);