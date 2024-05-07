CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE public.products (
	id uuid primary key not NULL DEFAULT uuid_generate_v4(),
	name_product varchar(255) NOT NULL,
	description text null,
	image text NULL,
	category varchar null,
	created_at timestamp not null DEFAULT now(),
	updated_at timestamp
);
