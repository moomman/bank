-- name: GetAccountById :one
SELECT * FROM account
WHERE id = ? ;

-- name: GetAccountForUpdate :one
select *
from account
where id = ?
for update;

-- name: CreateAccount :execresult
insert into account(owner, balance, currency,created_at)
value(?,?,?,?);

-- name: UpdateAccountById :exec
update account
set balance = balance + ?
where id = ?;

-- name: DeleteAccountById :exec
delete from account
where id = ?;
