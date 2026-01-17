create table skill_categories (
    skill_id int not null references skills(id) on delete cascade,
    category_id int not null references categories(id) on delete cascade,
    primary key (skill_id, category_id)
);
