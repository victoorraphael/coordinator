CREATE TABLE IF NOT EXISTS address (
  id SERIAL,
  uuid CHAR(36),
  street VARCHAR(255),
  city VARCHAR(255),
  zip varchar(8) not null,
  number int,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS subjects (
  id SERIAL,
  uuid CHAR(36),
  nome VARCHAR(255),

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS classroom (
  id SERIAL,
  uuid CHAR(36),
  name VARCHAR(100),

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS subjects_classroom (
  subject_id INT NOT NULL,
  classroom_id INT NOT NULL,
  FOREIGN KEY (subject_id) REFERENCES subjects (id),
  FOREIGN KEY (classroom_id) REFERENCES classroom (id),
  CONSTRAINT CLASSROOM_SUBJECT UNIQUE (classroom_id, subject_id)
);

CREATE TABLE IF NOT EXISTS school (
  id SERIAL,
  uuid CHAR(36),
  name varchar(255),
  address_id int,

  PRIMARY KEY (id),
  FOREIGN KEY (address_id) REFERENCES address (id)
);

CREATE TABLE IF NOT EXISTS school_classroom (
  school_id int,
  classroom_id int,
  FOREIGN KEY (school_id) REFERENCES school (id),
  FOREIGN KEY (classroom_id) REFERENCES classroom (id),
  CONSTRAINT SCHOOL_CLASSROOM_PK UNIQUE (school_id, classroom_id)
);

CREATE TABLE IF NOT EXISTS persons(
  id SERIAL PRIMARY KEY,
  uuid CHAR(36),
  name varchar(255) not null,
  email varchar(255) unique not null,
  phone varchar(11),
  birthdate timestamp not null,
  created_at timestamp default now(),
  type smallint,
  address_id int not null,

  CONSTRAINT addr_id FOREIGN KEY (address_id) REFERENCES address (id)
);

CREATE TABLE IF NOT EXISTS professor (
  id int not null,
  school_id int not null,

  FOREIGN KEY (id) REFERENCES persons (id),
  FOREIGN KEY (school_id) REFERENCES school (id),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS professor_subject (
  professor_id INT NOT NULL,
  subject_id INT NOT NULL,

  FOREIGN KEY (professor_id) REFERENCES professor (id),
  FOREIGN KEY (subject_id) REFERENCES subjects (id),
  PRIMARY KEY (professor_id, subject_id),
  CONSTRAINT PROFESSOR_SUBJECT_CT UNIQUE (professor_id, subject_id)
);

CREATE TABLE IF NOT EXISTS students (
  id int not null,
  classroom_id int not null,
  school_id int not null,

  FOREIGN KEY (id) REFERENCES persons (id),
  FOREIGN KEY (classroom_id) REFERENCES classroom (id),
  FOREIGN KEY (school_id) REFERENCES school (id),
  PRIMARY KEY (id)
);
