CREATE TABLE public.user
(
    id          SERIAL PRIMARY KEY,
    email       VARCHAR(100) NOT NULL,
    displayName varchar(50)  not null
);

CREATE TABLE public.family
(
    id    serial  not null primary key,
    name  text    not null,
    owner integer not null references public.user (id)
);

create table public.family_user
(
    family_id integer not null references public.family (id),
    user_id   integer not null references public.user (id)
);
create
index family_user_family_id on public.family_user (family_id);
create
index family_user_user_id on public.family_user (user_id);

create table public.ingredient
(
    id      serial primary key,
    name    text not null,
    si_unit text not null
);

CREATE TABLE public.recipe
(
    id        SERIAL  NOT NULL PRIMARY KEY,
    name      TEXT    NOT NULL,
    url       TEXT,
    family_id integer not null references public.family (id)
);

create table public.recipe_ingredient
(
    recipe_id     integer not null references public.recipe (id),
    ingredient_id integer not null references public.ingredient (id),
    amount        float   not null
);
create
index recipe_ingredient_recipe_id on public.recipe_ingredient (recipe_id);
create
index recipe_ingredient_ingredient_id on public.recipe_ingredient (ingredient_id);

CREATE TABLE public.session
(
    id        SERIAL      NOT NULL PRIMARY KEY,
    userId    INTEGER     NOT NULL,
    sessionId VARCHAR(64) NOT NULL,
    validTo   TIMESTAMP   NOT NULL
);
