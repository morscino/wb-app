-- +goose Up
-- +goose StatementBegin
create table loan_instalments
(
	id uuid
		constraint loan_instalments_pk
			primary key DEFAULT uuid_generate_v4(),
	user_id uuid not null,
    loan_id uuid not null,
    repayment_duration bigint not null,
	repayment_amount NUMERIC(14,2) not null,
    repayment_status varchar(100) not null,
    repayment_date timestamp not null,
    created_at timestamp default current_timestamp not null,
	updated_at timestamp default null,
	deleted_at timestamp default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table loan_instalments
-- +goose StatementEnd

