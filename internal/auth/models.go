package auth

import "context"

type User struct {
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, id int) error
	GetUserByEmail(ctx context.Context, email string) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}

type AuthUseCase interface {
	Register(ctx context.Context, user *User)
	Authenticate(ctx context.Context, email, password string) (*User, error)
	Authorize(ctx context.Context, user *User, permission string) bool
	GrantPermission(ctx context.Context, user *User, permission string) error
	RevokePermission(ctx context.Context, user *User, permission string) error
}

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpUserRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
	SignInUserRequest
}
