package repositories_test

import (
	"api/repositories"
	"api/types"
	"api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthRepoTestSuite struct {
	suite.Suite
}

func (suite *AuthRepoTestSuite) TestLogin() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Login when user exist", func(t *testing.T) {
		// Given
		const email = "testeur1@test.fr"

		expectedUser := types.User{
			ID:        1,
			FirstName: "testeur1",
			LastName:  "testeur1",
			Email:     "testeur1@test.fr",
			Password:  "testeur1",
			Role:      "retailer",
		}

		// When
		user, err := repositories.Login(utils.DbTest, email)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedUser, user, "User is not correct")
	})

	suite.T().Run("Login when user does not exist", func(t *testing.T) {
		// Given
		const email = "ceci_est_un_faux_mail@mail.fr"

		// When
		_, err := repositories.Login(utils.DbTest, email)

		// Then
		assert.NotNil(suite.T(), err, "Error is nil")
	})
}

func (suite *AuthRepoTestSuite) TestRegister() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Register when user does not exist", func(t *testing.T) {
		// Given
		userToRegister := types.User{
			FirstName: "testeur4",
			LastName:  "testeur4",
			Email:     "testeur4@test.fr",
			Password:  "testeur4",
			Role:      "customer",
		}

		expectedLastInsertId := int64(4)

		// When
		lastInsertId, err := repositories.Register(utils.DbTest, &userToRegister)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedLastInsertId, lastInsertId, "Last insert id is not correct")
	})
}

func TestAuthRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRepoTestSuite))
}
