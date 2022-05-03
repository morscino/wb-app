-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table loans
(
	id uuid
		constraint loans_pk
			primary key DEFAULT uuid_generate_v4(),
	user_id uuid not null,
    repayment_duration bigint not null,
	other_loans_amount NUMERIC(14,2) not null,
	loan_amount NUMERIC(14,2)  not null,
    account_number varchar(100) not null,
    account_name varchar(225) not null,
	bank varchar(100) not null,
    amount_paid NUMERIC(14,2) default null,
    balance NUMERIC(14,2) default null,
    status varchar(100) default null,
    loan_approval_date timestamp default null,
    repayment_status varchar(100) default null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null,
	deleted_at timestamp default null
);


create unique index loans_user_id_uindex
	on loans (user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table loans;
-- +goose StatementEnd
