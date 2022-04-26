CREATE TABLE restores (
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    id varchar(72),
    name varchar(72),
    namespace varchar(120),
    backup_name varchar(250),
    storage_secret varchar(120)
);