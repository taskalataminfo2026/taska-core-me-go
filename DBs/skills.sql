create table skills (
    id serial primary key,
    name varchar(75) not null,
    slug varchar(90) not null unique,
    is_active boolean not null default true,
    description text,
    created_at  timestamp(0),
    updated_at  timestamp(0)
);
