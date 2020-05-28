-- +goose Up
CREATE table books (
    id              serial primary key,
    title           text,
    description     text,
    meta            jsonb,
    created_at      timestamptz not null default now(),
    updated_at      timestamptz
);

INSERT INTO books (title, description, meta, updated_at)
VALUES
    ('Мастер и Маргарита', 'test description 1', '{}', now()),
    ('Граф Монте-Кристо', 'test description 2', null, null),
    ('Марсианин', 'test description 3', '{"author": "Энди Вейер"}', now());

-- +goose Down
drop table books;
