package repositories

import (
	"api/database"
	"api/types"
	"database/sql"
	"fmt"
)

func Login(db *sql.DB, email string) (types.User, error) {
	res := database.DoQueryRow(db, "SELECT * FROM users WHERE email = ?", email)

	user := new(types.User)
	errScan := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if errScan != nil {
		fmt.Printf("Error while scanning user: %s\n", errScan)
		if errScan == sql.ErrNoRows {
			fmt.Printf("There is no user with this email: %s\n", errScan)
			return types.User{}, errScan
		}
		return types.User{}, errScan
	}

	return *user, nil
}

func Register(db *sql.DB, user *types.User) (int64, error) {
	res, err := database.DoExec(db, "INSERT INTO users (firstName, lastName, email, password, role) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.Role)
	if err != nil {
		fmt.Printf("Error while creating user: %s\n", err)
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Error while getting user last inserted id: %s\n", err)
		return 0, err
	}

	return lastId, nil
}
