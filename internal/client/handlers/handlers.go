package handlers

import (
	"InternService/internal/auth"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AppContext struct {
	app     *fiber.App
	authUse auth.AuthUseCase
	//userUse user.UserUseCase
}

func InitHandlers(client *AppContext) {
	auth := client.app.Group("/api/auth")
	//group.Post("/signup", signUp)
	auth.Post("/signin", client.LogInHandler)
	//account := app.Group("/api/account")
	//account.Get("/", middleware.Authorize, getAccount)
	//group.Post("/avatar", middleware.Authorize, updateAvatar)
}

func (a *AppContext) LogInHandler(ctx *fiber.Ctx) error {
	var body auth.SignInUserRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}
	if err != nil {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	user, err := a.authUse.Authenticate()
	if err != nil {
		return err
	} else {

	}
	if email == "" || password == "" {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	trimmedEmail := strings.TrimSpace(email)
	trimmedPassword := strings.TrimSpace(password)
	if trimmedEmail == "" || trimmedPassword == "" {
		return response.Response(response.ResponseParams{
			Ctx:    ctx,
			Info:   constants.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
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

//func SignUpHandler(ctx *fiber.Ctx) error {
//	// check data
//	var body SignUpUserRequest
//	bodyParsingError := ctx.BodyParser(&body)
//	if bodyParsingError != nil {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//	email := body.Email
//	name := body.Name
//	password := body.Password
//	role := body.Role
//	if email == "" || name == "" || password == "" || role == "" {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.MissingData,
//			Status: fiber.StatusBadRequest,
//		})
//	}
//	trimmedEmail := strings.TrimSpace(email)
//	trimmedName := strings.TrimSpace(name)
//	trimmedPassword := strings.TrimSpace(password)
//	trimmedRole := strings.TrimSpace(role)
//	if trimmedEmail == "" || trimmedName == "" ||
//		trimmedPassword == "" || trimmedRole == "" {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.MissingData,
//			Status: fiber.StatusBadRequest,
//		})
//	}
//
//	// make sure that the role is correct
//	roles := utilities.Values(configuration.Roles)
//	if !utilities.IncludesString(roles, trimmedRole) {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.InvalidData,
//			Status: fiber.StatusBadRequest,
//		})
//	}
//
//	// load User schema
//	UserCollection := Instance.Database.Collection("User")
//
//	// check if email is already in use
//	existingRecord := UserCollection.FindOne(
//		ctx.Context(),
//		bson.D{{Key: "email", Value: trimmedEmail}},
//	)
//	existingUser := &User{}
//	existingRecord.Decode(existingUser)
//	if existingUser.ID != "" {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.EmailAlreadyInUse,
//			Status: fiber.StatusBadRequest,
//		})
//	}
//
//	// create a new User record, insert it and get back the ID
//	now := utilities.MakeTimestamp()
//	NewUser := new(User)
//	NewUser.Created = now
//	NewUser.Email = trimmedEmail
//	NewUser.ID = ""
//	NewUser.Name = trimmedName
//	NewUser.Role = trimmedRole
//	NewUser.Updated = now
//	insertionResult, insertionError := UserCollection.InsertOne(ctx.Context(), NewUser)
//	if insertionError != nil {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//	createdRecord := UserCollection.FindOne(
//		ctx.Context(),
//		bson.D{{Key: "_id", Value: insertionResult.InsertedID}},
//	)
//	createdUser := &User{}
//	createdRecord.Decode(createdUser)
//
//	// load Password schema
//	PasswordCollection := Instance.Database.Collection("Password")
//
//	// create password hash
//	hash, hashError := utilities.MakeHash(trimmedPassword)
//	if hashError != nil {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//
//	// create a new Password record and insert it
//	NewPassword := new(Password)
//	NewPassword.Created = now
//	NewPassword.Hash = hash
//	NewPassword.ID = ""
//	NewPassword.Updated = now
//	NewPassword.UserId = createdUser.ID
//	_, insertionError = PasswordCollection.InsertOne(ctx.Context(), NewPassword)
//	if insertionError != nil {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//
//	accessExpiration, expirationError := strconv.Atoi(os.Getenv("TOKENS_ACCESS_EXPIRATION"))
//	if expirationError != nil {
//		accessExpiration = 24
//	}
//	token, tokenError := utilities.GenerateJWT(utilities.GenerateJWTParams{
//		ExpiresIn: int64(accessExpiration),
//		UserId:    createdUser.ID,
//	})
//	if tokenError != nil {
//		return utilities.Response(utilities.ResponseParams{
//			Ctx:    ctx,
//			Info:   configuration.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//
//	return utilities.Response(utilities.ResponseParams{
//		Ctx: ctx,
//		Data: fiber.Map{
//			"token": token,
//			"user":  createdUser,
//		},
//	})
//}
