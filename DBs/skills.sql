create table skills (
    id serial primary key,                                      -- Identidad interna de la habilidad
    name varchar(75) not null,                                  -- Nombre humano de la habilidad
    slug varchar(90) not null unique,                           -- Identificador estable de sistema
    description text,                                           -- Alcance funcional de la habilidad
    avg_price_estimate float,                                   -- Referencia interna de precio típico
    requires_verification boolean default false,                -- Exige validación formal
    risk_level smallint default 1,                              -- Nivel de riesgo operativo/legal
    is_active boolean not null default true,                    -- Habilidad activa en el marketplace
    created_at timestamp(0)                                     -- Fecha de creación
);

