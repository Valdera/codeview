package service

import (
	"codeview/internal/dto/request"
	"context"
)

//go:generate mockery --name=AuthService --case underscore --testonly
type AuthService interface {
	Login(ctx context.Context, req *request.Login) (string, error)
	Register(ctx context.Context, req *request.Register) (string, error)
}
