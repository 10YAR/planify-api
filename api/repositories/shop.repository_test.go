package repositories_test

import (
	"api/repositories"
	"api/types"
	"api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ShopRepoTestSuite struct {
	suite.Suite
}

func (suite *ShopRepoTestSuite) TestGetShop() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)
	// Given
	ShopInfos := types.ShopInfos{
		ShopName:    "Dentest",
		Description: "dentiste test",
		Address:     "1 rue des dentistes, 00000 TestCity",
		PhoneNumber: "00 00 00 00 00",
	}
	expectedShop := types.Shop{
		ID:        1,
		ShopInfos: ShopInfos,
		CreatedAt: "2023-02-03 16:02:34",
		UserId:    1,
	}
	const shopId = "1"

	// When
	shop, _ := repositories.GetShop(utils.DbTest, shopId)

	// Then
	assert.Equal(suite.T(), expectedShop, shop, "Shop is not correct")
}

func (suite *ShopRepoTestSuite) TestCreateShop() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)
	// Given
	ShopInfos := types.ShopInfos{
		ShopName:    "ShopTest",
		Description: "shop test",
		Address:     "1 rue test, 00000 TestCity",
		PhoneNumber: "00 00 00 00 00",
	}
	shopToCreate := types.Shop{
		ShopInfos: ShopInfos,
		CreatedAt: "2023-03-01 11:00:00",
		UserId:    1,
	}
	expectedLastId := int64(3)

	// When
	lastId, _ := repositories.CreateShop(utils.DbTest, shopToCreate)

	// Then
	assert.Equal(suite.T(), expectedLastId, lastId, "Error while creating shop")
}

func TestShopRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ShopRepoTestSuite))
}
