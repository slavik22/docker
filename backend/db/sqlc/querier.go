// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateTutorial(ctx context.Context, arg CreateTutorialParams) (Tutorial, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteTutorial(ctx context.Context, id int32) error
	GetTutorial(ctx context.Context, id int32) (Tutorial, error)
	GetTutorials(ctx context.Context) ([]Tutorial, error)
	GetTutorialsByUser(ctx context.Context, userID int32) ([]Tutorial, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	UpdateTutorial(ctx context.Context, arg UpdateTutorialParams) (Tutorial, error)
}

var _ Querier = (*Queries)(nil)
