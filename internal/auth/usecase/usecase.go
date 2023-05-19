package usecase

import (
	"InternService/internal/auth"
	"context"
)

type AuthUseCase struct {
}

func (a AuthUseCase) Register(ctx context.Context, user *auth.User) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Authenticate(ctx context.Context, email, password string) (*auth.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Authorize(ctx context.Context, user *auth.User, permission string) bool {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) GrantPermission(ctx context.Context, user *auth.User, permission string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) RevokePermission(ctx context.Context, user *auth.User, permission string) error {
	//TODO implement me
	panic("implement me")
}

func NewUseCase(ctx context.Context) auth.AuthUseCase {
	return &AuthUseCase{}
}
