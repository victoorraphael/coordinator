CREATE TABLE IF NOT EXISTS school (
    id SERIAL,
    name varchar(255),
    address_id int,
    classroom_id int,

    PRIMARY KEY (id),
    FOREIGN KEY (address_id) REFERENCES address (id),
    FOREIGN KEY (classroom_id) REFERENCES classroom (id)
);