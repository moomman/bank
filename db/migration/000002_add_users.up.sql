create table user (
    id bigint primary key auto_increment,
    username varchar(255) not null,
    hash_password varchar(255) not null,
    full_name varchar(255) not null,
    email varchar(255) not null unique,
    create_time datetime not null,
    password_change_time datetime default '1000-01-01 00:00:00'
);

alter table account drop index idx_owner;
alter table account add unique index idx_owner_currency(owner,currency);