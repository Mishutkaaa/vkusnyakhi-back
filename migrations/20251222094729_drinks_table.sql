-- +goose Up
-- +goose StatementBegin
create table drinks(
id int not null,
name text,
image text,
categories int[],
brand int,
primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table drinks;
-- +goose StatementEnd
