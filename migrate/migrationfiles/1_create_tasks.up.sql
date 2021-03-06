create table
if not exists tasks
(
    id bigint unsigned auto_increment not null comment 'primary key' primary key,
    name varchar
(32) not null comment 'task name',
    code varchar
(64) not null unique comment 'task code, unique',
    type varchar
(32) not null comment 'task type, e.g. DELAY_JOB, DELAY_QUEUE',
    status varchar
(32) not null comment 'DISABLED->ENABLED',
    expired_at bigint unsigned not null comment 'current task expired time',
    cron varchar
(32) default null comment 'cron expressions, refer https://cron.qqe2.com',
    timeout int not null comment 'task execute timeout, unit is seconds',
    scheduling_category varchar
(32) not null comment 'scheduling category, e.g. singleton, multiple',
    executor varchar
(32) not null comment 'task execute method when expored, e.g. HTTP, RPC',
    deleted_at bigint unsigned not null default '0' comment 'deleted time',
    created_at bigint unsigned not null comment 'created time',
    created_by varchar
(32) default null comment 'created by',
    updated_at bigint unsigned not null comment 'last updated time',
    updated_by varchar
(32) default null comment 'last updated by'
) comment 'tasks' charset = utf8mb4;