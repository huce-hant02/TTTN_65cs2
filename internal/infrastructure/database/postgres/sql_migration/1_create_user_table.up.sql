drop table if exists users;
create table users
(
    id         bigserial
        primary key,
    username       varchar(255) not null unique,
    password       text not null ,
    status         varchar(255) not null default 'active',
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone null,
    deleted_at timestamp with time zone null,
);

create index idx_user_username
    on users (username);