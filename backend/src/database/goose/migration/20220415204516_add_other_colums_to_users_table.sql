-- +goose Up
-- +goose StatementBegin
alter table users
    add column if not exists user_type int2 default null,
    add column if not exists association_id uuid default null,
    add column if not exists association_branch varchar(100) default null,
    add column if not exists business_registration_date timestamp default null,
    add column if not exists business_rc_number varchar(100) default null,
    add column if not exists business_name varchar(100) default null,
    add column if not exists occupation varchar(100) default null,
    add column if not exists salary_range varchar(100) default null,
    add column if not exists date_of_birth timestamp default null,
    add column if not exists marital_status varchar(100) default null,
    add column if not exists means_of_identification int2 default null,
    add column if not exists profile_picture_url varchar(256) default null,
    add column if not exists telephone varchar(256) default null,
    add column if not exists document_url varchar(256) default null
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users
    drop column if exists user_type,
    drop column if exists association_id,
    drop column if exists association_branch,

    drop column if exists business_rc_number,
    drop column if exists occupation,
    drop column if exists salary_range,

    drop column if exists date_of_birth,
    drop column if exists marital_status,
    drop column if exists means_of_identification,

    drop column if exists profile_picture_url,
    drop column if exists document_url,

    drop column if exists business_registration_date
;
-- +goose StatementEnd


	
		