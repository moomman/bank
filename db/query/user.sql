-- name: CreateUser :execresult
insert into user(username, hash_password, full_name, email, create_time)
value(?,?,?,?,?);

-- name: GetUserByName :one
select * from user
where username = ?;

