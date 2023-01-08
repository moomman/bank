drop table user;
alter table account add index idx_owner(owner);
alter table account drop  index idx_owner_currency;