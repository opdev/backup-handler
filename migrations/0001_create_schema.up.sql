CREATE TABLE backups (
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    id varchar(72),
    name varchar(72),
    namespace varchar(60),
    state varchar(10),
    pod varchar(60),
    container varchar(60),
    command varchar(256),
    backup_location varchar(256),
    storage_secret varchar(128),
    kube_resource varchar
);