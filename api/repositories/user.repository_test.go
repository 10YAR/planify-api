package repositories_test

import (
	"api/repositories"
	"api/types"
	"api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserRepoTestSuite struct {
	suite.Suite
}

//func (suite *UserRepoTestSuite) TestGetUsers() {
//	pool, resource := utils.IntegrationTestSetup()
//	defer utils.IntegrationTestTeardown(pool, resource)
//
//	suite.T().Run("Get all users", func(t *testing.T) {
//		// Given
//		expectedUsers := []types.User{
//			{
//				ID:        1,
//				FirstName: "testeur1",
//				LastName:  "testeur1",
//				Email:     "",
//			}
//		}
//	}
//}

func (suite *UserRepoTestSuite) TestGetUser() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get a user", func(t *testing.T) {
		// Given
		expectedUser := types.User{
			ID:        1,
			FirstName: "testeur1",
			LastName:  "testeur1",
			Email:     "testeur1@test.fr",
			Password:  "testeur1",
			Role:      "retailer",
		}
		const userId = "1"

		// When
		user, err := repositories.GetUser(utils.DbTest, userId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedUser, user, "User is not correct")
	})
}

func (suite *UserRepoTestSuite) TestUpdateUser() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Update a user", func(t *testing.T) {
		// Given
		userToUpdate := types.User{
			ID:        1,
			FirstName: "testeur1",
			LastName:  "testeur1",
			Email:     "testeur@mail.fr",
			Password:  "testeur4",
			Role:      "customer",
		}
		const userId = "1"
		expectedRowsAffected := int64(1)

		// When
		rowsAffected, err := repositories.UpdateUser(utils.DbTest, &userToUpdate, userId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Rows affected is not correct")
	})
}

func (suite *UserRepoTestSuite) TestDeleteUser() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Delete a user without appointments or shops", func(t *testing.T) {
		// Given
		const userId = "3"
		expectedRowsAffected := int64(1)

		// When
		rowsAffected, err := repositories.DeleteUser(utils.DbTest, userId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Rows affected is not correct")
	})

	suite.T().Run("Delete a user with appointments or shops", func(t *testing.T) {
		// Given
		const userId = "2"

		// When
		_, err := repositories.DeleteUser(utils.DbTest, userId)

		// Then
		assert.NotNil(suite.T(), err, "Error is not nil")
	})
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
