#存放数据库脚本
create table account (
                         id bigint not null auto_increment primary key,
                         owner varchar(255) not null ,
                         balance bigint not null default 0,
                         currency varchar(255) not null default 'USD',
                         created_at datetime not null default now(),

    #constraint con_owner unique (owner),
                         unique index idx_owner(owner),
                         index idx_currency(currency),
                         index idx_create(created_at)
);

#代表了一对多的关系
create table entries (
                         id bigint not null auto_increment primary key,
                         account_id bigint not null,
                         amount bigint not null default 0,
                         created_at datetime not null default now(),

                         index idx_create(created_at)
);

create table transfers (
                           id bigint not null auto_increment primary key,
                           to_account_id bigint not null,
                           from_account_id bigint not null,
                           amount bigint not null default 0,
                           created_at datetime not null default now(),

                           index idx_create(created_at),
                           index idx_to_account(to_account_id),
                           index idx_from_account(from_account_id)
)