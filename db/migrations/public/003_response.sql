CREATE TABLE IF NOT EXISTS public."response"
(
    id serial NOT NULL,
    code int NOT NULL,
    message text NOT NULL,
    headers jsonb,
    body_raw bytea,
    body_text text,
    CONSTRAINT response_pkey PRIMARY KEY (id)
)

---- create above / drop below ----

DROP TABLE IF EXISTS public."response";
 