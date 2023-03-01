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
		shop := repositories.GetShop(utils.DbTest, shopId)

		// Then
		assert.Equal(suite.T(), expectedShop, shop, "Shop is not correct")
	})
}

func (suite *ShopRepoTestSuite) TestCreateShop() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Create a shop", func(t *testing.T) {
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
		expectedLastId := int64(4)

		// When
		lastId := repositories.CreateShop(utils.DbTest, &shopToCreate)

		// Then
		assert.Equal(suite.T(), expectedLastId, lastId, "Last id is not correct")
	})
}

//func (suite *ShopRepoTestSuite) TestGetShops() {
//	pool, resource := utils.IntegrationTestSetup()
//	defer utils.IntegrationTestTeardown(pool, resource)
//	// Given
//	ShopInfos1 := types.ShopInfos{
//		ShopName:    "Dentest",
//		Description: "dentiste test",
//		Address:     "1 rue des dentistes, 00000 TestCity",
//		PhoneNumber: "00 00 00 00 00",
//	}
//	expectedShop1 := types.Shop{
//		ID:        1,
//		ShopInfos: ShopInfos1,
//		CreatedAt: "2023-02-03 16:02:34",
//		UserId:    1,
//	}
//	expectedShops := []types.Shop{expectedShop1}
//
//	// When
//	shops, _ := repositories.GetShops()
//
//	// Then
//	assert.Equal(suite.T(), expectedShops, shops, "Shops are not correct")
//}

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
		rowsAffected := repositories.UpdateShop(utils.DbTest, &shopToUpdate, ShopId)

		// Then
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
		rowsAffected, _ := repositories.DeleteShop(utils.DbTest, shopId)

		// Then
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Rows affected is not correct")
	})

	//suite.T().Run("Shop with appointments", func(t *testing.T) {
	//	pool, resource := utils.IntegrationTestSetup()
	//	defer utils.IntegrationTestTeardown(pool, resource)
	//	// Given
	//	const shopId = 1
	//	expectedOutput := "Error 1451 (23000): Cannot delete or update a parent row: a foreign key constraint fails (`mysql`.`appointments`, CONSTRAINT `appointments_ibfk_1` FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`))"
	//
	//	// When
	//	_, err := repositories.DeleteShop(utils.DbTest, shopId)
	//	fmt.Printf("err: %v", err)
	//
	//	// Then
	//	assert.Equal(suite.T(), expectedOutput, err, "Rows affected is not correct")
	//})
}

//func (suite *ShopRepoTestSuite) TestGetShopAppointments() {
//	pool, resource := utils.IntegrationTestSetup()
//	defer utils.IntegrationTestTeardown(pool, resource)
//
//	suite.T().Run("Get a shop appointments", func(t *testing.T) {
//		// Given
//		const shopId = "1"
//
//		expectedAppointmentDateTimeInfos1 := types.AppointmentDateTimeInfos{
//			AppointmentDate:     "2023-03-03",
//			AppointmentTime:     "9:30:00",
//			AppointmentDateTime: "2023-03-03 10:30:00",
//		}
//		expectedAppointment1 := types.Appointment{
//			ID:                       1,
//			CustomerName:             "testeur_sans_compte_1",
//			AppointmentDateTimeInfos: expectedAppointmentDateTimeInfos1,
//			ShopId:                   1,
//		}
//
//		expectedAppointmentDateTimeInfos2 := types.AppointmentDateTimeInfos{
//			AppointmentDate:     "2023-03-03",
//			AppointmentTime:     "10:30:00",
//			AppointmentDateTime: "2023-03-03 10:30:00",
//		}
//		expectedAppointment2 := types.Appointment{
//			ID:                       1,
//			CustomerName:             "testeur3 testeur3",
//			AppointmentDateTimeInfos: expectedAppointmentDateTimeInfos2,
//			ShopId:                   1,
//		}
//
//		expectedAppointmentDateTimeInfos3 := types.AppointmentDateTimeInfos{
//			AppointmentDate:     "2023-03-03",
//			AppointmentTime:     "12:30:00",
//			AppointmentDateTime: "2023-03-03 10:30:00",
//		}
//		expectedAppointment3 := types.Appointment{
//			ID:                       1,
//			CustomerName:             "test_sans_compte_2",
//			AppointmentDateTimeInfos: expectedAppointmentDateTimeInfos3,
//			ShopId:                   1,
//		}
//
//		expectedAppointments := []types.Appointment{
//			expectedAppointment1,
//			expectedAppointment2,
//			expectedAppointment3,
//		}
//
//		// When
//		appointments := repositories.GetShopAppointments(utils.DbTest, shopId)
//
//		// Then
//		assert.Equal(suite.T(), expectedAppointments, appointments, "Appointments are not correct")
//	})
//}

func TestShopRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ShopRepoTestSuite))
}
