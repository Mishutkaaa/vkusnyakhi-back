-- +goose Up
-- +goose StatementBegin
create table categories(
id int not null,
name text,
primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table categories;
-- +goose StatementEnd
