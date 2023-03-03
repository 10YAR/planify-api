package repositories

import (
	"api/database"
	"api/types"
	"database/sql"
	"fmt"
)

func GetAppointments(db *sql.DB) ([]types.Appointment, error) {
	res, err := database.DoQuery(db, "SELECT * FROM appointments")

	if err != nil {
		fmt.Printf("Error while getting appointments: %s\n", err)
		return []types.Appointment{}, err
	}

	var appointments []types.Appointment
	for res.Next() {
		var appointment types.Appointment
		err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.AppointmentDateTime, &appointment.ShopId)
		if err != nil {
			fmt.Printf("Error while getting appointments: %s\n", err)
			return []types.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}
	if len(appointments) == 0 {
		fmt.Printf("No appointments: %s\n", sql.ErrNoRows)
		return []types.Appointment{}, sql.ErrNoRows
	}
	return appointments, nil
}

func GetUserAppointments(db *sql.DB, userId string) ([]types.Appointment, error) {
	res, err := database.DoQuery(db, "SELECT * FROM appointments WHERE user_id = ?", userId)

	if err != nil {
		fmt.Printf("Error while getting appointments: %s\n", err)
		return []types.Appointment{}, err
	}

	var appointments []types.Appointment
	for res.Next() {
		var appointment types.Appointment
		err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.AppointmentDateTime, &appointment.ShopId, &appointment.UserId, &appointment.Email)
		if err != nil {
			fmt.Printf("Error while getting appointments: %s\n", err)
			return []types.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}
	if len(appointments) == 0 {
		fmt.Printf("No appointments: %s\n", sql.ErrNoRows)
		return []types.Appointment{}, sql.ErrNoRows
	}
	return appointments, nil
}

func GetAppointment(db *sql.DB, id string) (types.Appointment, error) {
	res := database.DoQueryRow(db, "SELECT * FROM appointments WHERE id = ?", id)

	var appointment types.Appointment
	err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.AppointmentDateTime, &appointment.ShopId)
	if err != nil {
		fmt.Printf("Error while getting appointment: %s\n", err)
		if err == sql.ErrNoRows {
			fmt.Printf("There is no appointment with this id: %s\n", err)
			return types.Appointment{}, err
		}
		return types.Appointment{}, err
	}
	return appointment, nil
}

func CreateAppointment(db *sql.DB, appointment *types.Appointment) (int64, error) {
	res, err := database.DoExec(db, "INSERT INTO appointments (customer_name, appointment_date, appointment_time, appointment_date_time, shop_id, user_id, user_email) VALUES (?, ?, ?, CONCAT(?, ' ', ?), ?, ?, ?)", appointment.CustomerName, appointment.AppointmentDate, appointment.AppointmentTime, appointment.AppointmentDate, appointment.AppointmentTime, appointment.ShopId, appointment.UserId, appointment.Email)
	if err != nil {
		fmt.Printf("Error while creating appointment: %s\n", err)
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Error while getting last inserted id: %s\n", err)
		return 0, err
	}

	return lastId, nil
}

func UpdateAppointment(db *sql.DB, appointment *types.Appointment, id string) (int64, error) {
	res, err := database.DoExec(db, "UPDATE appointments SET customer_name = ?, appointment_date = ?, appointment_time = ?, appointment_date_time = CONCAT(?, ' ', ?), shop_id = ? WHERE id = ?", appointment.CustomerName, appointment.AppointmentDate, appointment.AppointmentTime, appointment.AppointmentDate, appointment.AppointmentTime, appointment.ShopId, id)
	if err != nil {
		fmt.Printf("Error while updating appointment: %s\n", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while getting rows affected: %s\n", err)
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteAppointment(db *sql.DB, id string) (int64, error) {
	res, err := database.DoExec(db, "DELETE FROM appointments WHERE id = ?", id)
	if err != nil {
		fmt.Printf("Error while deleting appointment: %s\n", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while getting rows affected: %s\n", err)
		return 0, err
	}

	return rowsAffected, nil
}
