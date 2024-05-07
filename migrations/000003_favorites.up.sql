 CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
 CREATE TABLE public.favorites(
 	id uuid primary key not NULL DEFAULT uuid_generate_v4(),
    product_id uuid NOT NULL,
    user_id uuid NOT NULL,
 	created_at timestamp not null DEFAULT now(),
 	CONSTRAINT product_fk FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE on update cascade,
 	CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE on update cascade
 );
 