CREATE TABLE public.family
(
    id            serial primary key,
    name          text not null,
    invitation_id text
);

CREATE TABLE public.user
(
    id           text PRIMARY KEY,
    email        VARCHAR(100) NOT NULL,
    display_name varchar(50)  not null,
    family_id    integer references public.family (id),
    created_at   timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE public.recipe
(
    id          SERIAL PRIMARY KEY,
    name        TEXT    NOT NULL,
    url         TEXT,
    family_id   integer not null references public.family (id),
    ingredients jsonb   not null
);
