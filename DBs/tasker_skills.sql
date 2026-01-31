create table tasker_skills (
    user_id int not null references users(id) on delete cascade,    -- Tasker dueño de la habilidad
    skill_id int not null references skills(id) on delete restrict, -- Habilidad que posee
    level smallint default 1,                                       -- Nivel de dominio
    years_experience smallint,                                      -- Años de experiencia
    jobs_completed int default 0,                                   -- Trabajos reales completados
    rating_avg numeric(2,1),                                        -- Promedio de calificación
    is_verified boolean default false,                              -- Validación por la plataforma
    last_used_at timestamp,                                         -- Último uso real
    created_at timestamp not null default now(),                    -- Registro de la habilidad
    primary key (user_id, skill_id)                                 -- Evita duplicar skills
);