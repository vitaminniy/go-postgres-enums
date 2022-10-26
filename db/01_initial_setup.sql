CREATE TYPE color AS ENUM (
	'red',
	'green',
	'blue',
	'white',
	'black'
);

CREATE TABLE lights (
  id    SERIAL PRIMARY KEY,
  color color NOT NULL
);
