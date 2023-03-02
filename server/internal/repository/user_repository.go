package repository

import (
	"codeview/internal/entity"
	"context"
)

//go:generate mockery --name=UserRepository --case underscore --testonly
type UserRepository interface {
	CreateUser(ctx context.Context, body *entity.User) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}
