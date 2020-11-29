create table 
if not exists scheduling_records
(
id bigint unsigned auto_increment not null comment 'primary key' primary key,
task_id bigint unsigned not null comment 'relation of task',
status varchar(32) not null comment 'scheduling result, task staus, lifecycle: new->runnable->running->finished/failed',
message text default null comment 'task execute result mask, e.g. success flag or failure reason or error log',
deleted_at bigint unsigned not null default '0' comment 'deleted time',
created_at bigint unsigned not null comment 'created time',
created_by varchar(32) default null comment 'created by',
updated_at bigint unsigned not null comment 'last updated time',
updated_by varchar(32) default null comment 'last updated by'
) comment 'task excute record' charset = utf8mb4; 