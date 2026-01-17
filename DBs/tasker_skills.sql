create table tasker_skills (
    user_id int not null references users(id) on delete cascade,
    skill_id int not null references skills(id) on delete restrict,
    level smallint default 1,
    is_verified boolean default false,
    created_at timestamp not null default now(),
    primary key (user_id, skill_id)
);
