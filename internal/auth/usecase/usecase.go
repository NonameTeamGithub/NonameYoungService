package usecase

import (
	"InternService/internal/auth"
	"InternService/internal/auth/repository"
	"InternService/internal/utilities"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	UserCollection, _ := a.rep.Mongo.GetCollection("Users")
	// check if email is already in use
	existingRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "email", Value: trimmedEmail}},
	)
	existingUser := &auth.User{}
	existingRecord.Decode(existingUser)
	if existingUser.ID != "" {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.EmailAlreadyInUse,
			Status: fiber.StatusBadRequest,
		})
	}
	now := utilities.MakeTimestamp()
	NewUser := new(auth.User)
	NewUser.Created = now
	NewUser.Email = trimmedEmail
	NewUser.ID = ""
	NewUser.Name = trimmedName
	NewUser.Role = trimmedRole
	NewUser.Updated = now
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
