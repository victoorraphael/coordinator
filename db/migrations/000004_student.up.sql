CREATE TABLE IF NOT EXISTS students (
    id int not null,
    FOREIGN KEY (id) REFERENCES persons (id),
    PRIMARY KEY (id)
);