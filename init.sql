CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
	id uuid primary key NOT NULL DEFAULT uuid_generate_v4(),
  	displayname varchar(100) NULL,
	first_name varchar(100) NULL ,
	last_name varchar(100) NULL,
    gender varchar(15) NULL,
	phone varchar(15) unique not NULL,
	email text unique NOT NULL,
	password text NOT NULL,
  	birth_date timestamp NULL,
	role varchar(6) NOT NULL DEFAULT 'user'::character varying,
	image text NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL
);

CREATE TABLE products (
	id uuid primary key not NULL DEFAULT uuid_generate_v4(),
	name_product varchar(255) NOT NULL,
	description text null,
	image text NULL,
	category varchar null,
	created_at timestamp not null DEFAULT now(),
	updated_at timestamp
);

 CREATE TABLE favorites(
 	id uuid primary key not NULL DEFAULT uuid_generate_v4(),
    product_id uuid NOT NULL,
    user_id uuid NOT NULL,
 	created_at timestamp not null DEFAULT now(),
 	CONSTRAINT product_fk FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE on update cascade,
 	CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE on update cascade
 );