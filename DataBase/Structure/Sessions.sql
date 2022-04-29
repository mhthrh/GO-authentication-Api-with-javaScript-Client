-- Table: voiceflex_schm.sessions

-- DROP TABLE IF EXISTS voiceflex_schm.sessions;

CREATE TABLE IF NOT EXISTS voiceflex_schm.sessions
(
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    session_id text COLLATE pg_catalog."default" NOT NULL,
    status boolean NOT NULL,
    valid_till timestamp without time zone NOT NULL,
    CONSTRAINT sessions_pkey PRIMARY KEY (id),
    CONSTRAINT user_frn FOREIGN KEY (user_id)
        REFERENCES voiceflex_schm.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS voiceflex_schm.sessions
    OWNER to postgres;