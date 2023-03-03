package repositories_test

import (
	"api/repositories"
	"api/types"
	"api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AppointmentRepoTestSuite struct {
	suite.Suite
}

func (suite *AppointmentRepoTestSuite) TestGetAppointments() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get all appointments", func(t *testing.T) {
		// Given

		// When

		// Then
	})
}

func (suite *AppointmentRepoTestSuite) TestGetUserAppointments() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get user appointmenta", func(t *testing.T) {
		// Given
		expectedAppointment := types.Appointment{
			ID:           2,
			CustomerName: "testeur3 testeur3",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "10:30:00",
				AppointmentDateTime: "2023-03-03 10:30:00",
			},
			ShopId: 1,
			UserId: 3,
			Email:  "testeur3@test.fr",
		}

		expectedAppointment1 := types.Appointment{
			ID:           5,
			CustomerName: "testeur3 testeur3",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "14:30:00",
				AppointmentDateTime: "2023-03-03 14:30:00",
			},
			ShopId: 2,
			UserId: 3,
			Email:  "testeur3@test.fr",
		}
		UserId := "3"

		expectedAppointments := []types.Appointment{expectedAppointment, expectedAppointment1}

		// When
		appointments, err := repositories.GetUserAppointments(utils.DbTest, UserId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedAppointments, appointments, "Appointments are not correct")
	})
}

func (suite *AppointmentRepoTestSuite) TestGetAppointment() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Get an appointment", func(t *testing.T) {
		// Given
		expectedAppointment := types.Appointment{
			ID:           1,
			CustomerName: "testeur_sans_compte_1",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "09:30:00",
				AppointmentDateTime: "2023-03-03 09:30:00",
			},
			ShopId: 1,
		}
		ShopId := "1"

		// When
		appointment, err := repositories.GetAppointment(utils.DbTest, ShopId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedAppointment, appointment, "Appointment is not correct")
	})
}

func (suite *AppointmentRepoTestSuite) TestCreateAppointment() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Create an appointment", func(t *testing.T) {
		// Given
		appointment := types.Appointment{
			CustomerName: "testeur3 testeur3",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate: "2023-03-03",
				AppointmentTime: "11:30:00",
			},
			ShopId: 1,
			UserId: 3,
			Email:  "testeur3@test.fr",
		}
		expectedLastId := int64(4)

		// When
		lastId, err := repositories.CreateAppointment(utils.DbTest, &appointment)

		// Then
		assert.Nil(t, err, "Error is not nil")
		assert.NotEqual(suite.T(), expectedLastId, lastId, "Appointment id is not correct")
	})
}

func (suite *AppointmentRepoTestSuite) TestUpdateAppointment() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Update an appointment", func(t *testing.T) {
		// Given
		appointment := types.Appointment{
			ID:           1,
			CustomerName: "testeur_sans_compte_1",
			AppointmentDateTimeInfos: types.AppointmentDateTimeInfos{
				AppointmentDate:     "2023-03-03",
				AppointmentTime:     "11:30:00",
				AppointmentDateTime: "2023-03-03 11:30:00",
			},
			ShopId: 1,
		}
		appointmentId := "1"
		expectedRowsAffected := int64(1)

		// When
		rowsAffected, err := repositories.UpdateAppointment(utils.DbTest, &appointment, appointmentId)

		// Then
		assert.Nil(suite.T(), err, "Error is not nil")
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Appointment is not correct")
	})
}

func (suite *AppointmentRepoTestSuite) TestDeleteAppointment() {
	pool, resource := utils.IntegrationTestSetup()
	defer utils.IntegrationTestTeardown(pool, resource)

	suite.T().Run("Delete an appointment", func(t *testing.T) {
		// Given
		appointmentId := "1"
		expectedRowsAffected := int64(1)

		// When
		rowsAffected, err := repositories.DeleteAppointment(utils.DbTest, appointmentId)

		// Then
		assert.Nil(t, err, "Error is not nil")
		assert.Equal(suite.T(), expectedRowsAffected, rowsAffected, "Appointment is not correct")
	})
}

func TestAppointmentsRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AppointmentRepoTestSuite))
}
