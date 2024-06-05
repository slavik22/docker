// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
)

type Tutorial struct {
	ID       int32  `json:"id"`
	UserID   int32  `json:"user_id"`
	Material string `json:"material"`
	Title    string `json:"title"`
}

type User struct {
	ID             int32        `json:"id"`
	Email          string       `json:"email"`
	HashedPassword string       `json:"hashed_password"`
	Name           string       `json:"name"`
	IsAdmin        sql.NullBool `json:"is_admin"`
}
