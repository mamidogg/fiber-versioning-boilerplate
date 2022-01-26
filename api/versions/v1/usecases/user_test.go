package usecases

import (
	"testing"

	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/entities"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/repositories/mocks"
	"github.com/stretchr/testify/suite"
)

type userUsecaseSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the generated mocked version of our repository
	repository *mocks.IUserRepository
	// the functionalities we want to test
	usecase IUserUsecase
}

func (suite *userUsecaseSuite) SetupTest() {
	// instantiate the mocked version of repository
	repository := new(mocks.IUserRepository)

	// inject the repository to usecase, since usecase needs repository to work
	usecase := NewUserUsecase(repository)

	// assign them as the suite properties
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *userUsecaseSuite) TestGetAllUser_Positive() {
	var repoMockInfo entities.User
	suite.repository.On("GetAllUser", "1").Return(&repoMockInfo, nil)

	result, err := suite.usecase.GetAllUser()
	suite.Nil(err, "no error when GetAllUser")
	suite.Equal(repoMockInfo, *result, "result should be equal")
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}
