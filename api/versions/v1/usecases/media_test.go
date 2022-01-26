package usecases

import (
	"testing"

	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/repositories/mocks"
	"github.com/stretchr/testify/suite"
)

type mediaUsecaseSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the generated mocked version of our repository
	repository *mocks.MediaRepository
	// the functionalities we want to test
	usecase MediaUsecase
}

func (suite *mediaUsecaseSuite) SetupTest() {
	// instantiate the mocked version of repository
	repository := new(mocks.MediaRepository)

	// inject the repository to usecase, since usecase needs repository to work
	usecase := NewMediaUsecase(repository)

	// assign them as the suite properties
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *mediaUsecaseSuite) TestGetMediaByID_Exists_Positive() {
	id := "1"
	var repoMockInfo interface{}
	suite.repository.On("GetMediaByID", id).Return(&repoMockInfo, nil)

	result, err := suite.usecase.GetMediaByID(id)
	suite.Nil(err, "no error when return the tweet")
	suite.Equal(repoMockInfo, *result, "result and tweet should be equal")
}

func (suite *mediaUsecaseSuite) TestGetMediaByIDS_Exists_Positive() {
	ids := []string{
		"1",
		"2",
		"3",
		"4",
	}
	var repoMockInfo []interface{}
	suite.repository.On("GetMediaByIDS", ids).Return(&repoMockInfo, nil)

	result, err := suite.usecase.GetMediaByIDS(ids)
	suite.Nil(err, "no error when GetAllUser")
	suite.Equal(repoMockInfo, *result, "result should be equal")
}

func TestMediaUsecase(t *testing.T) {
	suite.Run(t, new(mediaUsecaseSuite))
}
