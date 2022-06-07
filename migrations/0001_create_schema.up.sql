CREATE TABLE backups (
    created_at varchar(32),
    updated_at  varchar(32),
    deleted_at  varchar(32),
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