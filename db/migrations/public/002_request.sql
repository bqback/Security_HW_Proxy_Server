CREATE TABLE IF NOT EXISTS public."request"
(
    id serial NOT NULL,
    method text NOT NULL,
    scheme text NOT NULL, 
    host text NOT NULL,
    path text NOT NULL,
    get jsonb,
    headers jsonb,
    cookies jsonb,
    post jsonb,
    CONSTRAINT request_pkey PRIMARY KEY (id)
)

---- create above / drop below ----

DROP TABLE IF EXISTS public."request";
