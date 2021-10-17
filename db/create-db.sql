-- Table: public.clipboardhistory

-- DROP TABLE IF EXISTS public.clipboardhistory;

CREATE TABLE IF NOT EXISTS public.clipboardhistory
(
    id integer NOT NULL DEFAULT nextval('clipboardhistory_id_seq'::regclass),
    insert_datetime timestamp with time zone,
    data character varying COLLATE pg_catalog."default",
    CONSTRAINT clipboardhistory_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.clipboardhistory
    OWNER to postgres;