package repositories

import (
	"api/database"
	"api/types"
	"database/sql"
	"fmt"
)

func GetShops() ([]types.Shop, error) {
	res, err := database.DoQuery("SELECT * FROM shops")

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

func CreateShop(db *sql.DB, shop *types.Shop) (int64, error) {
	res, err := database.DoExec(db, "INSERT INTO shops (shop_name, description, address, phone_number, created_at, user_id) VALUES (?, ?, ?, ?, ?, ?)", shop.ShopName, shop.Description, shop.Address, shop.PhoneNumber, shop.CreatedAt, shop.UserId)
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
	res, err := database.DoExec(db, "UPDATE shops SET shop_name = ?, description = ?, address = ?, phone_number = ?, created_at = ?, user_id = ? WHERE id = ?", shop.ShopName, shop.Description, shop.Address, shop.PhoneNumber, shop.CreatedAt, shop.UserId, id)
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

//func GetShopAppointments(db *sql.DB, id string) []types.AppointmentDateTimeInfos {
//	res, err := database.DoQuery(db, "SELECT appointment_date, appointment_time, appointment_date_time FROM appointments WHERE shop_id = ?", id)
//
//	if err != nil {
//		fmt.Printf("Error while getting shop appointments: %s\n", err)
//	}
//
//	var appointments []types.AppointmentDateTimeInfos
//	for res.Next() {
//		var appointment types.AppointmentDateTimeInfos
//		err := res.Scan(&appointment.AppointmentDate, &appointment.AppointmentTime, &appointment.AppointmentDateTime)
//		if err != nil {
//			fmt.Printf("Error while getting shop appointments: %s\n", err)
//		}
//
//		appointments = append(appointments, appointment)
//	}
//	if len(appointments) == 0 {
//		fmt.Printf("No appointments: %s\n", err)
//	}
//	return appointments
//}
