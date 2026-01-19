create table skill_categories (
    id serial primary key,                                   -- Identidad de la relación
    skill_id int not null,                                   -- Habilidad asociada
    category_id int not null,                                -- Categoría donde aplica
    is_primary boolean default false,                        -- Habilidad principal en esta categoría
    is_active boolean default true,                          -- Relación activa o retirada
    unique(skill_id, category_id)                            -- Evita duplicados
);
