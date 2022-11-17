CREATE TABLE IF NOT EXISTS persons(
  id uuid DEFAULT uuid_generate_v4 (),
  name varchar(255),
  email varchar(255),
  phone varchar(11),
  birthdate timestamp,
  created_at timestamp,

  PRIMARY KEY (id)
);