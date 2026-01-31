create table categories (
    id serial primary key,                                      -- Identidad interna del mercado
    root_id int references categories(id),                      -- Vertical principal (Belleza, Hogar, Reparaciones)
    parent_id int references categories(id),                    -- Categoría padre en la jerarquía
    name varchar(75) not null,                                  -- Nombre visible del mercado
    slug varchar(90) not null unique,                           -- Identificador estable (URLs, APIs, SEO)
    description text,                                           -- Define qué tipo de servicios viven aquí
    icon varchar(255),                                          -- Ícono representativo para UI
    is_active boolean not null default true,                    -- Mercado habilitado o cerrado
    sort_order int default 0,                                   -- Prioridad estratégica en listados
    created_at timestamp(0),                                    -- Fecha de creación (cohortes y madurez)
    updated_at timestamp(0),                                    -- Fecha de actualización (cohortes y madurez)
);
