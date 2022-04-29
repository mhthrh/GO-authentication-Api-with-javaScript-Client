-- Table: voiceflex_schm.users

-- DROP TABLE IF EXISTS voiceflex_schm.users;

CREATE TABLE IF NOT EXISTS voiceflex_schm.users
(
    id uuid NOT NULL,
    first_name text COLLATE pg_catalog."default" NOT NULL,
    last_name text COLLATE pg_catalog."default" NOT NULL,
    user_name text COLLATE pg_catalog."default" NOT NULL,
    pass text COLLATE pg_catalog."default" NOT NULL,
    cell_number text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    create_date date NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT user_unq UNIQUE (user_name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS voiceflex_schm.users
    OWNER to postgres;