CREATE TABLE IF NOT EXISTS public."request"
(
    id serial NOT NULL,
    method text NOT NULL,
    scheme text NOT NULL, 
    host text NOT NULL,
    path text NOT NULL,
    get json,
    headers json,
    cookies json,
    post json,
    body_raw bytea,
    body_text text,
    CONSTRAINT request_pkey PRIMARY KEY (id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS public."request";
