-- +goose Up
-- +goose StatementBegin
alter table users
    add column if not exists state varchar(256) default null,
    add column if not exists local_government varchar(256) default null,
    add column if not exists bvn varchar(100) default null
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users
    drop column if exists state,
    drop column if exists local_government,
    drop column if exists bvn
;
-- +goose StatementEnd


	
		