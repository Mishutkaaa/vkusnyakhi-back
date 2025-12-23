-- +goose Up
-- +goose StatementBegin
create table brand(
id serial not null,
name text,
primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table brand;
-- +goose StatementEnd
