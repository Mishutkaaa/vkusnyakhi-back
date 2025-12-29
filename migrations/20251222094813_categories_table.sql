-- +goose Up
-- +goose StatementBegin
create table categories(
id serial not null,
name text,
type text,
primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table categories;
-- +goose StatementEnd
