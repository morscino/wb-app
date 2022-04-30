-- +goose Up
-- +goose StatementBegin
create table if not exists associations
(
	id uuid
		constraint associations_pk
			primary key DEFAULT uuid_generate_v4(),
	name varchar(225) not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null,
	deleted_at timestamp default null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table associations;
-- +goose StatementEnd