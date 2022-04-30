-- +goose Up
-- +goose StatementBegin
create table if NOT exists waitlists
(
	id uuid
		constraint waitlists_pk
			primary key DEFAULT uuid_generate_v4(),
	full_name varchar(225) not null,
    telephone varchar(225) not null,
    business_name varchar(225) default null,
	mode smallint not null,
	email varchar(225) not null UNIQUE,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null,
	deleted_at timestamp default null
);

create unique index waitlists_email_uindex
	on waitlists (email);

create index waitlists_full_name_index
	on waitlists (full_name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table waitlists;
-- +goose StatementEnd

