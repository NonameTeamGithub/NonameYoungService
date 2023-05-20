package handlers

import (
	"InternService/internal/auth"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type AppContext struct {
	Log     zerolog.Logger
	App     *fiber.App
	AuthUse auth.AuthUseCase
	//userUse user.UserUseCase
}

func InitHandlers(client *AppContext) {
	auth := client.App.Group("/api/auth")
	auth.Post("/signup", client.SignUpHandler)
	auth.Post("/signin", client.LogInHandler)
	//account := app.Group("/api/account")
	//account.Get("/", middleware.Authorize, getAccount)
	//group.Post("/avatar", middleware.Authorize, updateAvatar)
}

func (a *AppContext) LogInHandler(ctx *fiber.Ctx) error {
	var body auth.SignInUserRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	user, token, err := a.AuthUse.Authenticate(ctx, body.Email, body.Password)
	if err != nil {
		return err
	}
	return response.Response(response.ResponseParams{
		Ctx: ctx,
		Data: fiber.Map{
			"check": "true",
			"token": token,
			"user":  user,
		},
	})
	//// load User schema
	//UserCollection := Instance.Database.Collection("User")
	//
	//// find a user
	//rawUserRecord := UserCollection.FindOne(
	//	ctx.Context(),
	//	bson.D{{Key: "email", Value: trimmedEmail}},
	//)
	//userRecord := &User{}
	//rawUserRecord.Decode(userRecord)
	//if userRecord.ID == "" {
	//	return response.Response(response.ResponseParams{
	//		Ctx:    ctx,
	//		Info:   constants.ResponseMessages.AccessDenied,
	//		Status: fiber.StatusUnauthorized,
	//	})
	//}
	//
	//// load Password schema
	//PasswordCollection := Instance.Database.Collection("Password")
	//
	//// find a password
	//rawPasswordRecord := PasswordCollection.FindOne(
	//	ctx.Context(),
	//	bson.D{{Key: "userId", Value: userRecord.ID}},
	//)
	//passwordRecord := &Password{}
	//rawPasswordRecord.Decode(passwordRecord)
	//if passwordRecord.ID == "" {
	//	return response.Response(response.ResponseParams{
	//		Ctx:    ctx,
	//		Info:   constants.ResponseMessages.AccessDenied,
	//		Status: fiber.StatusUnauthorized,
	//	})
	//}
	//
	//// compare hashes
	//passwordIsValid := utilities.CompareHashes(trimmedPassword, passwordRecord.Hash)
	//if !passwordIsValid {
	//	return response.Response(response.ResponseParams{
	//		Ctx:    ctx,
	//		Info:   constants.ResponseMessages.AccessDenied,
	//		Status: fiber.StatusUnauthorized,
	//	})
	//}
	//
	//accessExpiration, expirationError := strconv.Atoi(os.Getenv("TOKENS_ACCESS_EXPIRATION"))
	//if expirationError != nil {
	//	accessExpiration = 24
	//}
	//token, tokenError := jwtokens.GenerateJWT(jwtokens.GenerateJWTParams{
	//	ExpiresIn: int64(accessExpiration),
	//	UserId:    userRecord.ID,
	//})
	//if tokenError != nil {
	//	return response.Response(response.ResponseParams{
	//		Ctx:    ctx,
	//		Info:   constants.ResponseMessages.InternalServerError,
	//		Status: fiber.StatusInternalServerError,
	//	})
	//}
	// )
}

func (a *AppContext) SignUpHandler(ctx *fiber.Ctx) error {
	// check data
	var body auth.SignUpUserRequest
	bodyParsingError := ctx.BodyParser(&body)
	if bodyParsingError != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	fmt.Println(body)
	return a.AuthUse.Register(ctx, body)
}
