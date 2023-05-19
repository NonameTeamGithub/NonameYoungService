package usecase

import (
	"InternService/internal/auth"
	"InternService/internal/storage"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"context"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AuthUseCase struct {
	storage storage.Storage
}

func (a AuthUseCase) Register(ctx *fiber.Ctx, user *auth.User) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Authenticate(ctx *fiber.Ctx, email, password string) (*auth.User, string, error) {
	if email == "" || password == "" {
		return nil, "", response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	trimmedEmail := strings.TrimSpace(email)
	trimmedPassword := strings.TrimSpace(password)
	if trimmedEmail == "" || trimmedPassword == "" {
		return nil, "", response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	
}

func (a AuthUseCase) Authorize(ctx *fiber.Ctx, user *auth.User, permission string) bool {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) GrantPermission(ctx *fiber.Ctx, user *auth.User, permission string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) RevokePermission(ctx *fiber.Ctx, user *auth.User, permission string) error {
	//TODO implement me
	panic("implement me")
}

func NewUseCase(ctx context.Context) auth.AuthUseCase {
	return &AuthUseCase{}
}
