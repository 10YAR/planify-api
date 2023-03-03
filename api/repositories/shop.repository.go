package repositories

import (
	"api/database"
	"api/types"
	"database/sql"
	"fmt"
)

func GetShops(db *sql.DB) ([]types.Shop, error) {
	res, err := database.DoQuery(db, "SELECT * FROM shops")

	if err != nil {
		fmt.Printf("Error while getting shops from database: %s\n", err)
		return []types.Shop{}, err
	}

	var shops []types.Shop
	for res.Next() {
		var shop types.Shop
		err := res.Scan(&shop.ID, &shop.ShopName, &shop.Description, &shop.Address, &shop.PhoneNumber, &shop.CreatedAt, &shop.UserId)
		if err != nil {
			fmt.Printf("Error while scanning shops: %s\n", err)
			return []types.Shop{}, err
		}

		shops = append(shops, shop)
	}
	if len(shops) == 0 {
		fmt.Printf("No shops: %s\n", err)
		return []types.Shop{}, sql.ErrNoRows
	}
	return shops, nil
}

func GetShop(db *sql.DB, id string) (types.Shop, error) {
	res := database.DoQueryRow(db, "SELECT * FROM shops WHERE id = ?", id)

	var shop types.Shop
	err := res.Scan(&shop.ID, &shop.ShopName, &shop.Description, &shop.Address, &shop.PhoneNumber, &shop.CreatedAt, &shop.UserId)
	if err != nil {
		fmt.Printf("Error while getting shop: %s\n", err)
		if err == sql.ErrNoRows {
			fmt.Printf("There is no shop with this id: %s\n", err)
			return types.Shop{}, err
		}
		return types.Shop{}, err
	}
	return shop, nil
}

func GetShopAvailabilities(db *sql.DB, id string) ([]types.ShopAvailability, error) {
	res, err := database.DoQuery(db, "SELECT day_of_week, duration, start_time, end_time FROM shop_availability WHERE shop_id = ?", id)

	if err != nil {
		fmt.Printf("Error while getting shop availabilities from database: %s\n", err)
		return []types.ShopAvailability{}, err
	}

	var availabilities []types.ShopAvailability
	for res.Next() {
		var availability types.ShopAvailability
		err := res.Scan(&availability.DayOfWeek, &availability.Duration, &availability.StartTime, &availability.EndTime)
		if err != nil {
			fmt.Printf("Error while scanning shop availabilities: %s\n", err)
			return []types.ShopAvailability{}, err
		}

		availabilities = append(availabilities, availability)
	}
	if len(availabilities) == 0 {
		fmt.Printf("No availabilities: %s\n", err)
		return []types.ShopAvailability{}, sql.ErrNoRows
	}
	return availabilities, nil
}

func GetShopAppointments(db *sql.DB, id string) ([]types.Appointment, error) {
	res, err := database.DoQuery(db, "SELECT id, customer_name, appointment_date, appointment_time, appointment_date_time, shop_id FROM appointments WHERE shop_id = ?", id)

	if err != nil {
		fmt.Printf("Error while getting shop appointments from database: %s\n", err)
		return []types.Appointment{}, err
	}

	var appointments []types.Appointment
	for res.Next() {
		var appointment types.Appointment
		err := res.Scan(&appointment.ID, &appointment.CustomerName, &appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.AppointmentDateTime, &appointment.ShopId)
		if err != nil {
			fmt.Printf("Error while scanning shop appointments: %s\n", err)
			return []types.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}
	if len(appointments) == 0 {
		fmt.Printf("No appointments: %s\n", err)
		return []types.Appointment{}, sql.ErrNoRows
	}
	return appointments, nil
}

func GetShopsByUserId(db *sql.DB, id string) ([]types.Shop, error) {
	res, err := database.DoQuery(db, "SELECT * FROM shops WHERE user_id = ?", id)

	if err != nil {
		fmt.Printf("Error while getting shops: %s\n", err)
		return []types.Shop{}, err
	}

	var shops []types.Shop
	for res.Next() {
		var shop types.Shop
		err := res.Scan(&shop.ID, &shop.ShopName, &shop.Description, &shop.Address, &shop.PhoneNumber, &shop.CreatedAt, &shop.UserId)
		if err != nil {
			fmt.Printf("Error while getting shops: %s\n", err)
			return []types.Shop{}, err
		}

		shops = append(shops, shop)
	}
	if len(shops) == 0 {
		fmt.Printf("No shops: %s\n", err)
		return []types.Shop{}, sql.ErrNoRows
	}
	return shops, nil
}

func CreateShop(db *sql.DB, shopAvailabilities *types.ShopAvailabilities) (int64, error) {
	res, err := database.DoExec(db, "INSERT INTO shops (shop_name, description, address, phone_number, user_id) VALUES (?, ?, ?, ?, ?)", shopAvailabilities.ShopName, shopAvailabilities.Description, shopAvailabilities.Address, shopAvailabilities.PhoneNumber, shopAvailabilities.UserId)

	for _, availability := range shopAvailabilities.Availabilities {
		_, errShopAvailability := database.DoExec(db, "INSERT INTO shop_availability (day_of_week, duration, start_time, end_time, shop_id) VALUES (?, ?, ?, ?, ?)", availability.DayOfWeek, availability.Duration, availability.StartTime, availability.EndTime, availability.ShopId)
		if errShopAvailability != nil {
			fmt.Printf("Error while creating shop availability: %s\n", err)
			return 0, err
		}
	}
	if err != nil {
		fmt.Printf("Error while creating shop: %s\n", err)
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Error while getting last inserted id: %s\n", err)
		return 0, err
	}

	return lastId, nil
}

func UpdateShop(db *sql.DB, shop *types.Shop, id string) (int64, error) {
	res, err := database.DoExec(db, "UPDATE shops SET shop_name = ?, description = ?, address = ?, phone_number = ?, user_id = ? WHERE id = ?", shop.ShopName, shop.Description, shop.Address, shop.PhoneNumber, shop.UserId, id)
	if err != nil {
		fmt.Printf("Error while updating shop: %s\n", err)
		return 0, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while getting affected rows: %s\n", err)
		return 0, err
	}

	return affectedRows, nil
}

func DeleteShop(db *sql.DB, id string) (int64, error) {
	res, err := database.DoExec(db, "DELETE s FROM shops s LEFT JOIN appointments a ON s.id = a.shop_id WHERE s.id = ?", id)

	if err != nil {
		fmt.Printf("Error while deleting shop: %s\n", err)
		return 0, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while getting affected rows: %s\n", err)
		return 0, err
	}

	return affectedRows, nil
}
