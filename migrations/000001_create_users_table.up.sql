CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public."group"
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    title text NOT NULL,
    descriptions text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.contact
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username text NOT NULL,
    given_name text NOT NULL,
    email text NOT NULL,
    birthday_date date NOT NULL,
    group_id UUID REFERENCES public.group(id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.phone
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    country_code INT NOT NULL,
    operator INT NOT NULL,
    "number" BIGINT NOT NULL,
    contact_id UUID REFERENCES public.contact(id),
    PRIMARY KEY (id)
);


