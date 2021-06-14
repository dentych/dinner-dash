CREATE TABLE public.family
(
    id            serial primary key,
    name          text not null,
    invitation_id text
);

CREATE TABLE public.user
(
    username        text PRIMARY KEY,
    email           VARCHAR(100) NOT NULL,
    hashed_password text         not null,
    family_id       integer references public.family (id),
    created_at      timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE public.recipe
(
    id          SERIAL PRIMARY KEY,
    name        TEXT    NOT NULL,
    url         TEXT,
    family_id   integer not null references public.family (id),
    description text    not null,
    ingredients jsonb   not null
);
