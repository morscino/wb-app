-- +goose Up
-- +goose StatementBegin
ALTER Table users
    drop column if exists salary_range
    ;

ALTER Table users
   add column if not exists salary NUMERIC(14,2) default null;
     
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER Table users
    add COLUMN salary_range VARCHAR(100);    
-- +goose StatementEnd
