-- +goose Up
-- +goose StatementBegin
create table brand(
id int not null,
name text,
primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table brand;
-- +goose StatementEnd
