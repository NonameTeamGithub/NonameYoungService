package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, id int) error
	GetUserByEmail(ctx context.Context, email string) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}

type AuthUseCase interface {
	Register(ctx *fiber.Ctx, body SignUpUserRequest) error
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

type User struct {
	AvatarLink string `json:"avatarLink"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Created    int64  `json:"created"`
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated    int64  `json:"updated"`
}

type Password struct {
	Hash    string `json:"hash"`
	UserId  string `json:"userId" bson:"userId"`
	Created int64  `json:"created"`
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated int64  `json:"updated"`
}
