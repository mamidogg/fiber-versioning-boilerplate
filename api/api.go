package api

import (
	"fmt"
	"log"

	"github.com/aofdev/fiber-versioning-boilerplate/api/adapters"
	handlers "github.com/aofdev/fiber-versioning-boilerplate/api/handlers"
	"github.com/aofdev/fiber-versioning-boilerplate/api/utilities"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/repositories"
	routesV1 "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/routes"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/usecases"

	_ "github.com/aofdev/fiber-versioning-boilerplate/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Fiber Versioning Boilerplate
// @version 1.0
// @description This is a sample swagger for Fiber

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func StartServer() {
	// Connect to the mysql database
	mysqlAdapter := adapters.NewMysqlServer(adapters.ENV{
		utilities.ViperEnvVariable("MYSQL_HOST"),
		utilities.ViperEnvVariable("MYSQL_PORT"),
		utilities.ViperEnvVariable("MYSQL_DATABASE_NAME"),
		utilities.ViperEnvVariable("MYSQL_USERNAME"),
		utilities.ViperEnvVariable("MYSQL_PASSWORD"),
	})

	// Connect to the mongo database
	mongoAdapter := adapters.NewMongoAdapter(
		utilities.ViperEnvVariable("MONGO_URI"),
		utilities.ViperEnvVariable("MONGO_DATABASE"),
		utilities.ViperEnvVariable("MONGO_COLLECTION"),
	)
	defer mongoAdapter.CloseMongoAdapter()

	// Initialize repository and usecase
	initMediaRepoV1 := repositories.NewMediaRepository(mongoAdapter)
	initMediaUsecaseV1 := usecases.NewMediaUsecase(initMediaRepoV1)

	initUserRepoV1 := repositories.NewUserRepository(mysqlAdapter)
	initUserUsecaseV1 := usecases.NewUserUsecase(initUserRepoV1)

	// Initialize fiber
	app := fiber.New(fiber.Config{
		//Prefork:      true,
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/docs/*", swagger.Handler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello fiber versioning boilerplate")
	})

	v1 := app.Group("/v1")
	routesV1.MediaRouter(v1, initMediaUsecaseV1)
	routesV1.UserRouter(v1, initUserUsecaseV1)

	port := fmt.Sprintf(":%s", utilities.ViperEnvVariable("APP_PORT"))
	log.Fatal(app.Listen(port))
}
