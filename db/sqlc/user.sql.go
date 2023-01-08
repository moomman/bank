// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
insert into user(username, hash_password, full_name, email, create_time)
value(?,?,?,?,?)
`

type CreateUserParams struct {
	Username     string    `json:"username"`
	HashPassword string    `json:"hash_password"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	CreateTime   time.Time `json:"create_time"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.HashPassword,
		arg.FullName,
		arg.Email,
		arg.CreateTime,
	)
}

const getUserByName = `-- name: GetUserByName :one
select id, username, hash_password, full_name, email, create_time, password_change_time from user
where username = ?
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashPassword,
		&i.FullName,
		&i.Email,
		&i.CreateTime,
		&i.PasswordChangeTime,
	)
	return &i, err
}