CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE public.users (
	id uuid primary key NOT NULL DEFAULT uuid_generate_v4(),
  	displayname varchar(100) NULL,
	first_name varchar(100) NULL ,
	last_name varchar(100) NULL,
    gender varchar(15) NULL,
	phone varchar(15) not NULL,
	email text NOT NULL,
	password text NOT NULL,
  	birth_date timestamp NULL,
	role varchar(6) NOT NULL DEFAULT 'user'::character varying,
	image text NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL
);