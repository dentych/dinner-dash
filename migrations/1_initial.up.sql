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

create table public.ingredient
(
    id      serial primary key,
    name    text not null,
    si_unit text not null
);

CREATE TABLE public.recipe
(
    id        SERIAL PRIMARY KEY,
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
