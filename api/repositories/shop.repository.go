package repositories

import (
	"api/database"
	"api/types"
	"api/utils"
	"database/sql"
)

func GetShops() ([]types.Shop, types.HttpResponse) {
	res, err := database.DoQuery("SELECT * FROM shops")

	var errorMessage types.HttpResponse
	if err != nil {
		errorMessage = utils.E503("Error while getting shops", err)
	}

	var shops []types.Shop
	for res.Next() {
		var shop types.Shop
		err := res.Scan(&shop.ID, &shop.ShopName, &shop.Description, &shop.Address, &shop.PhoneNumber, &shop.CreatedAt, &shop.UserId)
		if err != nil {
			errorMessage = utils.E503("Error while getting shops", err)
		}

		shops = append(shops, shop)
	}
	if len(shops) == 0 {
		errorMessage = utils.E503("No shops", err)
	}
	return shops, errorMessage
}

func GetShop(db *sql.DB, id string) (types.Shop, types.HttpResponse) {
	res := database.DoQueryRow(db, "SELECT * FROM shops WHERE id = ?", id)

	var errorMessage types.HttpResponse
	var shop types.Shop
	err := res.Scan(&shop.ID, &shop.ShopName, &shop.Description, &shop.Address, &shop.PhoneNumber, &shop.CreatedAt, &shop.UserId)
	if err != nil {
		errorMessage = utils.E503("Error while getting shop", err)
		if err == sql.ErrNoRows {
			errorMessage = utils.E503("There is no shop with this id.", err)
		}
	}
	return shop, errorMessage
}

func CreateShop(db *sql.DB, shop types.Shop) (int64, types.HttpResponse) {
	var errorMessage types.HttpResponse
	res, err := database.DoExec(db, "INSERT INTO shops (shop_name, description, address, phone_number, created_at, user_id) VALUES (?, ?, ?, ?, ?, ?)", shop.ShopName, shop.Description, shop.Address, shop.PhoneNumber, shop.CreatedAt, shop.UserId)
	if err != nil {
		errorMessage = utils.E503("Error while creating shop", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		errorMessage = utils.E503("Error while getting last inserted id", err)
	}

	return lastId, errorMessage
}
