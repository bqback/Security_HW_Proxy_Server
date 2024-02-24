CREATE TABLE IF NOT EXISTS public.request_response
(
    id_request serial NOT NULL,
    id_response serial NOT NULL,
    CONSTRAINT request_response_pkey PRIMARY KEY (id_request, id_response),
    CONSTRAINT request_response_id_request_fkey FOREIGN KEY (id_request)
        REFERENCES public."request" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT request_response_id_response_fkey FOREIGN KEY (id_response)
        REFERENCES public."response" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

---- create above / drop below ----

DROP TABLE IF EXISTS public.request_response;