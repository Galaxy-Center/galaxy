create table
if not exists task_configs
(
id bigint unsigned auto_increment not null comment 'primary key' primary key,
headers JSON default null comment 'config headers info, e.g. auth secret',
content JSON default null comment 'config task execute need content',
deleted_at bigint unsigned not null default '0' comment 'deleted time',
created_at bigint unsigned not null comment 'created time',
created_by varchar
(32) default null comment 'created by',
updated_at bigint unsigned not null comment 'last updated time',
updated_by varchar
(32) default null comment 'last updated by'
) comment 'task excute configuration' charset = utf8mb4;