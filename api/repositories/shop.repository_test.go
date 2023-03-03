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

func (suite *ShopRepoTestSuite) TestGetShops() {

	suite.T().Run("Get the shops of user", func(t *testing.T) {
		pool, resource := utils.IntegrationTestSetup()
		defer utils.IntegrationTestTeardown(pool, resource)
		// Given
		expectedShop1 := types.Shop{
			ID: 1,
			ShopInfos: types.ShopInfos{
				ShopName:    "Dentest",
				Description: "dentiste test",
				Address:     "1 rue des dentistes, 00000 TestCity",
				PhoneNumber: "00 00 00 00 00",
			},
			CreatedAt: "2023-02-03 16:02:34",
			UserId:    1,
		}
		expectedShop2 := types.Shop{
			ID: 2,
			ShopInfos: types.ShopInfos{
				ShopName:    "orltest",
				Description: "ORL test",
				Address:     "1 rue des orls, 00000 TestCity",
				PhoneNumber: "00 00 00 00 01",
			},
			CreatedAt: "2023-02-28 11:00:00",
			UserId:    2,
		}
		expectedShop3 := types.Shop{
			ID: 3,
			ShopInfos: types.ShopInfos{
				ShopName:    "shop3",
				Description: "shop sans appointments",
				Address:     "1 rue des shop3, 00000 TestCity",
				PhoneNumber: "00 00 00 00 02",
			},
			CreatedAt: "2023-02-28 11:00:00",
			UserId:    2,
		}

		expectedShops := []types.Shop{expectedShop1, expectedShop2, expectedShop3}

		// When
		shops, _ := repositories.GetShops(utils.DbTest)

		// Then
		assert.Equal(suite.T(), expectedShops, shops, "Shops are not correct")
	})
}

func (suite *ShopRepoTestSuite) TestGetShop() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get a shop", func(t *testing.T) {
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
		shop, err := repositories.GetShop(utils.DbTest, shopId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedShop, shop, "Shop is not correct")
	})
}

func (suite *ShopRepoTestSuite) TestGetShopAvailabilities() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get a shop availabilities", func(t *testing.T) {
		// Given
		const shopId = "1"

		expectedAvailability1 := types.ShopAvailability{
			DayOfWeek: "tuesday",
			Duration:  15,
			StartTime: "09:00:00",
			EndTime:   "19:00:00",
		}
		expectedAvailability2 := types.ShopAvailability{
			DayOfWeek: "wednesday",
			Duration:  15,
			StartTime: "09:00:00",
			EndTime:   "19:00:00",
		}
		expectedAvailability3 := types.ShopAvailability{
			DayOfWeek: "thursday",
			Duration:  15,
			StartTime: "09:00:00",
			EndTime:   "17:00:00",
		}
		expectedAvailability4 := types.ShopAvailability{
			DayOfWeek: "friday",
			Duration:  15,
			StartTime: "09:00:00",
			EndTime:   "17:00:00",
		}
		expectedAvailability5 := types.ShopAvailability{
			DayOfWeek: "saturday",
			Duration:  15,
			StartTime: "09:00:00",
			EndTime:   "19:00:00",
		}

		expectedAvailabilities := []types.ShopAvailability{
			expectedAvailability1,
			expectedAvailability2,
			expectedAvailability3,
			expectedAvailability4,
			expectedAvailability5,
		}

		// When
		shopAvailabilities, err := repositories.GetShopAvailabilities(utils.DbTest, shopId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedAvailabilities, shopAvailabilities, "Availabilities are not correct")
	})
}

func (suite *ShopRepoTestSuite) TestGetShopAppointments() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get a shop appointments", func(t *testing.T) {
		// Given
		const shopId = "1"

		expectedAppointment1 := types.Appointment{
			ID:           1,
			CustomerName: "testeur_sans_compte_1",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "09:30:00",
				AppointmentDateTime: "2023-03-03 09:30:00",
			},
			ShopId: 1,
		}

		expectedAppointment2 := types.Appointment{
			ID:           2,
			CustomerName: "testeur3 testeur3",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "10:30:00",
				AppointmentDateTime: "2023-03-03 10:30:00",
			},
			ShopId: 1,
		}

		expectedAppointment3 := types.Appointment{
			ID:           3,
			CustomerName: "test_sans_compte_2",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "12:30:00",
				AppointmentDateTime: "2023-03-03 12:30:00",
			},
			ShopId: 1,
		}

		expectedAppointments := []types.Appointment{
			expectedAppointment1,
			expectedAppointment2,
			expectedAppointment3,
		}

		// When
		shopAppointments, err := repositories.GetShopAppointments(utils.DbTest, shopId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedAppointments, shopAppointments, "Appointments are not correct")
	})
}

func (suite *ShopRepoTestSuite) TestCreateShop() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Create a shop", func(t *testing.T) {
		// Given
		shopToCreate := types.ShopAvailabilities{
			Shop: types.Shop{
				ShopInfos: types.ShopInfos{
					ShopName:    "ShopTest",
					Description: "shop test",
					Address:     "1 rue test, 00000 TestCity",
					PhoneNumber: "00 00 00 00 00",
				},
				CreatedAt: "2023-03-01 11:00:00",
				UserId:    1,
			},
			Availabilities: []types.ShopAvailabilityWithShopId{
				{
					ShopAvailability: types.ShopAvailability{
						DayOfWeek: "monday",
						Duration:  15,
						StartTime: "09:00:00",
						EndTime:   "19:00:00",
					},
					ShopId: 1,
				},
			},
		}
		expectedLastId := int64(4)

		// When
		lastId, err := repositories.CreateShop(utils.DbTest, &shopToCreate)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedLastId, lastId, "Last id is not correct")
	})
}

func (suite *ShopRepoTestSuite) TestGetShopsByUserId() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get shops by user id", func(t *testing.T) {
		// Given
		ShopInfos1 := types.ShopInfos{
			ShopName:    "Dentest",
			Description: "dentiste test",
			Address:     "1 rue des dentistes, 00000 TestCity",
			PhoneNumber: "00 00 00 00 00",
		}
		expectedShop1 := types.Shop{
			ID:        1,
			ShopInfos: ShopInfos1,
			CreatedAt: "2023-02-03 16:02:34",
			UserId:    1,
		}

		expectedShops := []types.Shop{expectedShop1}
		const userId = "1"

		// When
		shops, err := repositories.GetShopsByUserId(utils.DbTest, userId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedShops, shops, "Shops are not correct")
	})
}

func (suite *ShopRepoTestSuite) TestUpdateShop() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Update a shop", func(t *testing.T) {
		// Given
		ShopInfos := types.ShopInfos{
			ShopName:    "ShopTest",
			Description: "dentiste test",
			Address:     "1 rue des dentistes, 00000 TestCity",
			PhoneNumber: "00 00 00 00 00",
		}
		shopToUpdate := types.Shop{
			ID:        1,
			ShopInfos: ShopInfos,
			CreatedAt: "2023-02-03 16:02:34",
			UserId:    1,
		}
		const ShopId = "1"
		expectedRowsAffected := int64(1)

		// When
		rowsAffected, err := repositories.UpdateShop(utils.DbTest, &shopToUpdate, ShopId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Rows affected is not correct")
	})
}

func (suite *ShopRepoTestSuite) TestDeleteShop() {
	suite.T().Run("Shop without appointments", func(t *testing.T) {
		pool, resource := utils.IntegrationTestSetup()
		defer utils.IntegrationTestTeardown(pool, resource)

		// Given
		const shopId = "3"
		expectedRowsAffected := int64(1)

		// When
		rowsAffected, err := repositories.DeleteShop(utils.DbTest, shopId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Rows affected is not correct")
	})

	suite.T().Run("Shop with appointments", func(t *testing.T) {
		pool, resource := utils.IntegrationTestSetup()
		defer utils.IntegrationTestTeardown(pool, resource)
		// Given
		const shopId = "1"

		// When
		_, err := repositories.DeleteShop(utils.DbTest, shopId)

		// Then
		assert.NotNil(suite.T(), err, "It should return an error saying that the shop has appointments")
	})
}

func TestShopRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ShopRepoTestSuite))
}
