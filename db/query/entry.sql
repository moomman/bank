-- name: GetEntriesByAccountId :many
select * from entries
where account_id = ?;

-- name: CreateEntry :exec
insert into entries
(account_id, amount, created_at)
value (?,?,?);

-- name: UpdateEntryByAccountId :exec
update entries
set amount = ?
where account_id = ?;

-- name: deleteEntryByTime :exec
delete from entries
where created_at < ?;