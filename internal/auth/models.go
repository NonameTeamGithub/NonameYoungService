package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

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
	Register(ctx *fiber.Ctx, user *User)
	Authenticate(ctx *fiber.Ctx, email, password string) (*User, string, error)
	Authorize(ctx *fiber.Ctx, user *User, permission string) bool
	GrantPermission(ctx *fiber.Ctx, user *User, permission string) error
	RevokePermission(ctx *fiber.Ctx, user *User, permission string) error
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
