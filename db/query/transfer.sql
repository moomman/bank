-- name: GetTransfersByAccountId :many
select * from transfers
where (to_account_id or from_account_id) = ?;

-- name: CreateTransfer :exec
insert into transfers
(to_account_id, from_account_id, amount, created_at)
value (?,?,?,?);

-- name: UpdateTransferByAccountId :exec
update transfers
set amount = ?
where to_account_id = ?;

-- name: deleteTransferByTime :exec
delete from transfers
where created_at < ?;