CREATE TABLE IF NOT EXISTS school_classroom (
    school_id int,
    classroom_id int,
    FOREIGN KEY (school_id) REFERENCES school (id),
    FOREIGN KEY (classroom_id) REFERENCES classroom (id),
    CONSTRAINT SCHOOL_CLASSROOM_PK UNIQUE (school_id, classroom_id)
);