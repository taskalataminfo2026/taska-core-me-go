create table categories (
    id serial primary key,
    name varchar(75) not null,
    slug varchar(90) not null unique,
    parent_id int references categories(id),
    is_active boolean not null default true,
    description text,
    icon varchar(255),
    sort_order int default 0,
    created_at  timestamp(0),
    updated_at  timestamp(0)
);
