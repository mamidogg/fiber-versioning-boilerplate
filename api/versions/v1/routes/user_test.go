package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/aofdev/fiber-versioning-boilerplate/api/handlers"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/entities"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/usecases/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/suite"
)

type userRouterSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the mocked version of the usecase
	usecase *mocks.IUserUsecase
	// testing server to be used the handler
	testingServer    *fiber.App
	testingServerUrl string
}

func (suite *userRouterSuite) SetupSuite() {
	// create a mocked version of usecase
	usecase := new(mocks.IUserUsecase)

	// create and run the testing server
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})
	app.Use(cors.New())
	app.Use(recover.New())

	v1 := app.Group("/v1")
	UserRouter(v1, suite.usecase)

	app.Listen(":3000")

	testingServer := app

	testingServerUrl := "http://localhost:3000"

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.testingServerUrl = testingServerUrl
	suite.usecase = usecase
}

func (suite *userRouterSuite) TearDownSuite() {
	defer suite.testingServer.Shutdown()
}

func (suite *userRouterSuite) TestGetAllUser_Positive() {
	usecaseMockInfo := entities.User{
		ID:          "222",
		AccountCode: "1234",
		FirstName:   "Unit",
		LastName:    "Test",
		Email:       "unit.test@gmail.com",
	}
	suite.usecase.On("GetAllUser").Return(&usecaseMockInfo, nil)

	// calling the testing server given the provided request body
	url := fmt.Sprintf("%s/v1/users", suite.testingServerUrl)
	response, err := http.Get(url)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := entities.ResponseGetAllUser{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(responseBody.Status, "success")
}

func TestUserRouter(t *testing.T) {
	suite.Run(t, new(userRouterSuite))
}
