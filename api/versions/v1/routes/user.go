package routes

import (
	"github.com/aofdev/fiber-versioning-boilerplate/api/handlers"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/usecases"
	"github.com/gofiber/fiber/v2"
)

type IUserRouter interface {
	getAllUser(u usecases.UserUsecase) fiber.Handler
}

// UserRouter is group routers for user
func UserRouter(app fiber.Router, u usecases.IUserUsecase) {
	app.Get("/users", getAllUser(u))
}

// getAllUser
// @Summary Lists all user
// @Description get object user
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.HTTPSuccess{data=[]entities.ResponseGetAllUser}
// @Failure 400 {object} handlers.HTTPError
// @Failure 404 {object} handlers.HTTPError
// @Failure 500 {object} handlers.HTTPError
// @Router /v1/users [get]
func getAllUser(u usecases.IUserUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := u.GetAllUser()
		if err != nil {
			_ = c.JSON(&handlers.HTTPError{Status: "error", Error: err.Error()})
		}

		return c.JSON(&handlers.HTTPSuccess{Status: "success", Data: result})
	}
}
