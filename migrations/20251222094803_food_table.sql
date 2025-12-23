-- +goose Up
-- +goose StatementBegin
create table food(
id serial not null,
name text,
image text,
categories int[],
brand int,
primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table food;
-- +goose StatementEnd
