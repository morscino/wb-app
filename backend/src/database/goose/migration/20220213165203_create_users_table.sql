-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
	id uuid
		constraint users_pk
			primary key DEFAULT uuid_generate_v4(),
	first_name varchar(225) not null,
    last_name varchar(225) not null,
	email varchar(225) not null UNIQUE,
    salt varchar(225) not null,
	password text not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null,
	deleted_at timestamp default null
);

create unique index users_email_uindex
	on users (email);

create index users_first_name_index
	on users (first_name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
