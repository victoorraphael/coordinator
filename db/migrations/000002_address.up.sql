CREATE TABLE IF NOT EXISTS address (
    id SERIAL,
    street VARCHAR(255),
    city VARCHAR(255),
    zip varchar(8) not null,
    number int,

    PRIMARY KEY (id)
);