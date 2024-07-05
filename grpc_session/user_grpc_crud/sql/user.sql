CREATE TABLE public.users (
	id serial NOT NULL,
	name varchar NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	created_at timestamp DEFAULT NOW() NOT NULL,
	updated_at timestamp DEFAULT NOW() NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);
