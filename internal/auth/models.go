package auth

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	CreateUser(ctx *fiber.Ctx, user User, password string) (string, *User, error)
	GetUserById(ctx context.Context, id int) error
	GetUserByEmail(ctx *fiber.Ctx, email string) error
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
	Email      string `json:"email" bson:"email"`
	Name       string `json:"name" bson:"name"`
	Role       string `json:"role" bson:"role"`
	Created    int64  `json:"created" bson:"created"`
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated    int64  `json:"updated" bson:"updated"`
}

type Password struct {
	Hash    string `json:"hash" bson:"hash"`
	UserId  string `json:"userId" bson:"userId"`
	Created int64  `json:"created" bson:"created"`
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Updated int64  `json:"updated" bson:"updated"`
}
