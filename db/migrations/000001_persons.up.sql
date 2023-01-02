CREATE TABLE IF NOT EXISTS persons(
  id SERIAL,
  uuid uuid DEFAULT uuid_generate_v4(),
  name varchar(255) not null,
  email varchar(255) unique not null,
  phone varchar(11),
  birthdate timestamp not null,
  created_at timestamp default now(),
  type smallint,

  PRIMARY KEY (id)
);