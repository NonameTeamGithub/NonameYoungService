package usecase

import (
	"InternService/internal/auth"
	"InternService/internal/utilities"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"InternService/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"strings"
)

type AuthUseCase struct {
	log zerolog.Logger
	rep auth.AuthRepository
}

func (a AuthUseCase) Register(ctx *fiber.Ctx, body auth.SignUpUserRequest) error {
	// make sure that the role is correct
	email, name, password, role := body.Email, body.Name, body.Password, body.Role
	if email == "" || name == "" || password == "" || role == "" {
		a.log.Warn().Msg("a empty body")
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	trimmedEmail := strings.TrimSpace(email)
	trimmedName := strings.TrimSpace(name)
	trimmedPassword := strings.TrimSpace(password)
	trimmedRole := strings.TrimSpace(role)
	if trimmedEmail == "" || trimmedName == "" ||
		trimmedPassword == "" || trimmedRole == "" {
		a.log.Warn().Msg("something trimmed wrong")
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	roles := utilities.Values(constants.Roles)
	if !utilities.IncludesString(roles, trimmedRole) {
		a.log.Warn().Msg("someone string wrong")
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
	}
	err := a.rep.GetUserByEmail(ctx, email)
	if err != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
	}
	now := utilities.MakeTimestamp()
	NewUser := auth.User{
		Email:   trimmedEmail,
		Name:    trimmedName,
		Role:    trimmedRole,
		Created: now,
		Updated: now,
	}
	token, createdUser, err := a.rep.CreateUser(ctx, NewUser, trimmedPassword)
	if err != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
	}
	return response.Response(response.ResponseParams{
		Ctx: ctx,
		Data: fiber.Map{
			"token": token,
			"user":  createdUser,
		},
	})
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
	return &auth.User{}, "", nil
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

func NewUseCase(authR auth.AuthRepository) auth.AuthUseCase {
	return &AuthUseCase{logger.GetLogger(), authR}
}
