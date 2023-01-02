CREATE TABLE IF NOT EXISTS professor (
    id int not null,
    specialization int not null,

    FOREIGN KEY (id) REFERENCES persons (id),
    PRIMARY KEY (id)
);