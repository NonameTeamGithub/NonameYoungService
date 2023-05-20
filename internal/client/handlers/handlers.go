package handlers

import (
	"InternService/internal/auth"
	"InternService/internal/models"
	"InternService/internal/user/candidate"
	"InternService/internal/user/candidate/repository"
	"InternService/internal/utilities/constants"
	"InternService/internal/utilities/response"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type AppContext struct {
	App     *fiber.App
	AuthUse auth.AuthUseCase
	//CandidateUse candidate.UseCase
	//InternUse
	//Storage *repository.CandidateRepository
	Storage *repository.CandidateRepository
}

func InitHandlers(client *AppContext) {
	//candidate handlers
	candidates := client.App.Group("/users/candidates")
	candidates.Post("/", client.CreateCandidate)
	candidates.Get("/:id", client.GetCandidate)
	candidates.Delete("/:id", client.DeleteCandidate)
	candidates.Put("/:id", client.DeleteCandidate)
	//curators

	//
	auth := client.App.Group("/api/auth")
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

//func (a *AppContext) SignUpHandler(ctx *fiber.Ctx) error {
//	// check data
//	var body auth.SignUpUserRequest
//	bodyParsingError := ctx.BodyParser(&body)
//	if bodyParsingError != nil {
//		return response.Response(response.ResponseParams{
//			Ctx:    ctx,
//			Info:   constants.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//	a.authUse.Register(ctx, body)
//
//	// load User schema
//
//	// create a new User record, insert it and get back the ID
//	now := utilities.MakeTimestamp()
//	NewUser := new(auth.User)
//	NewUser.Created = now
//	NewUser.Email = trimmedEmail
//	NewUser.ID = ""
//	NewUser.Name = trimmedName
//	NewUser.Role = trimmedRole
//	NewUser.Updated = now
//	insertionResult, insertionError := UserCollection.InsertOne(ctx.Context(), NewUser)
//	if insertionError != nil {
//		return response.Response(response.ResponseParams{
//			Ctx:    ctx,
//			Info:   constants.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//	createdRecord := UserCollection.FindOne(
//		ctx.Context(),
//		bson.D{{Key: "_id", Value: insertionResult.InsertedID}},
//	)
//	createdUser := &auth.User{}
//	createdRecord.Decode(createdUser)
//
//	// load Password schema
//	PasswordCollection := Instance.Database.Collection("Password")
//
//	// create password hash
//	hash, hashError := utilities.MakeHash(trimmedPassword)
//	if hashError != nil {
//		return response.Response(response.ResponseParams{
//			Ctx:    ctx,
//			Info:   constants.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//
//	// create a new Password record and insert it
//	NewPassword := new(auth.Password)
//	NewPassword.Created = now
//	NewPassword.Hash = hash
//	NewPassword.ID = ""
//	NewPassword.Updated = now
//	NewPassword.UserId = createdUser.ID
//	_, insertionError = PasswordCollection.InsertOne(ctx.Context(), NewPassword)
//	if insertionError != nil {
//		return response.Response(response.ResponseParams{
//			Ctx:    ctx,
//			Info:   constants.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//
//	accessExpiration, expirationError := strconv.Atoi(os.Getenv("TOKENS_ACCESS_EXPIRATION"))
//	if expirationError != nil {
//		accessExpiration = 24
//	}
//	token, tokenError := jwtokens.GenerateJWT(jwtokens.GenerateJWTParams{
//		ExpiresIn: int64(accessExpiration),
//		UserId:    createdUser.ID,
//	})
//	if tokenError != nil {
//		return response.Response(response.ResponseParams{
//			Ctx:    ctx,
//			Info:   constants.ResponseMessages.InternalServerError,
//			Status: fiber.StatusInternalServerError,
//		})
//	}
//
//	return response.Response(response.ResponseParams{
//		Ctx: ctx,
//		Data: fiber.Map{
//			"token": token,
//			"user":  createdUser,
//		},
//	})
//}

func (a *AppContext) CreateCandidate(ctx *fiber.Ctx) error {
	m := models.CandidateModel{DB: a.Storage.DB}
	var user candidate.Candidate
	if err := ctx.BodyParser(&user); err != nil {
		log.Warn().Err(err)
		ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	if err := m.Create(user); err != nil {
		log.Warn().Err(err)
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
	}
	ctx.JSON(&fiber.Map{
		"success": true,
		"message": "User successfully created",
	})
	return nil
}
func (a *AppContext) GetCandidate(ctx *fiber.Ctx) error {
	m := models.CandidateModel{DB: a.Storage.DB}
	var user candidate.Candidate
	id := ctx.Params("id")
	if err := m.GetById(&user, id); err != nil {
		log.Warn().Err(err)
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	ctx.JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched candidate",
		"user":    user,
	})
	return nil
}
func (a *AppContext) DeleteCandidate(ctx *fiber.Ctx) error {
	m := models.CandidateModel{DB: a.Storage.DB}
	id := ctx.Params("id")
	if err := m.Delete(id); err != nil {
		log.Warn().Err(err)
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	ctx.JSON(&fiber.Map{
		"success": true,
		"message": "Successfully delete candidate",
	})
	return nil
}
