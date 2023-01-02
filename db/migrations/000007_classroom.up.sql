CREATE TABLE IF NOT EXISTS classroom (
     id SERIAL,
     name VARCHAR(100),
     subject_id INT,
     student_id INT,

     PRIMARY KEY (id),
     FOREIGN KEY (subject_id) REFERENCES subjects (id),
     FOREIGN KEY (student_id) REFERENCES students (id)
);