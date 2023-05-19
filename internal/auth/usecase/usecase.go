package usecase

import (
	"InternService/internal/auth"
	"InternService/internal/auth/repository"
	"InternService/internal/utilities"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/jwtokens"
	"InternService/internal/utilities/response"
	"context"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"strings"
)

type AuthUseCase struct {
	rep repository.Auth
}

func (a AuthUseCase) Register(ctx *fiber.Ctx, body auth.SignUpUserRequest) error {
	// make sure that the role is correct
	email, name, password, role := body.Email, body.Name, body.Password, body.Role
	if email == "" || name == "" || password == "" || role == "" {
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
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	roles := utilities.Values(constants.Roles)
	if !utilities.IncludesString(roles, trimmedRole) {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
	}
	err := a.rep.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	now := utilities.MakeTimestamp()
	NewUser := auth.User{
		AvatarLink: "",
		Email:      trimmedEmail,
		Name:       trimmedName,
		Role:       trimmedRole,
		Created:    now,
		ID:         "",
		Updated:    now,
	}
	err := a.rep.CreateUser(ctx, NewUser)
	PasswordCollection, _ := a.rep.Mongo.GetCollection("Password")

	// create password hash
	hash, hashError := utilities.MakeHash(trimmedPassword)
	if hashError != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	// create a new Password record and insert it
	NewPassword := new(auth.Password)
	NewPassword.Created = now
	NewPassword.Hash = hash
	NewPassword.ID = ""
	NewPassword.Updated = now
	NewPassword.UserId = createdUser.ID
	_, insertionError = PasswordCollection.InsertOne(ctx.Context(), NewPassword)
	if insertionError != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	accessExpiration, expirationError := strconv.Atoi(os.Getenv("TOKENS_ACCESS_EXPIRATION"))
	if expirationError != nil {
		accessExpiration = 24
	}
	token, tokenError := jwtokens.GenerateJWT(jwtokens.GenerateJWTParams{
		ExpiresIn: int64(accessExpiration),
		UserId:    createdUser.ID,
	})
	if tokenError != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
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
