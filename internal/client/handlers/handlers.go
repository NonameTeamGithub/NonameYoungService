package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func InitHandlers(app *fiber.App) {
	LogInHandler(app)
}

func LogInHandler(app *fiber.App) {
	app.Get("")
}

func SignUpHandler(app *fiber.App) {

}
