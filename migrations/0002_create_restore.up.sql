CREATE TABLE restores (
    created_at varchar(32),
    updated_at varchar(32),
    deleted_at varchar(32),
    id varchar(72),
    name varchar(72),
    namespace varchar(120),
    backup_name varchar(250),
    storage_secret varchar(120)
);