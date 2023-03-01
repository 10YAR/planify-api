package repositories

import (
	"api/database"
	"api/types"
	"database/sql"
	"fmt"
)

func GetUsers() ([]types.User, error) {
	rows, err := database.DoQuery("SELECT * FROM users")

	if err != nil {
		fmt.Printf("Error while getting users from database: %s\n", err)
		return []types.User{}, err
	}

	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			fmt.Printf("Error while scanning users: %s\n", err)
			return []types.User{}, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		fmt.Printf("No users found\n")
		return []types.User{}, sql.ErrNoRows
	}

	return users, nil
}

func GetUser(db *sql.DB, id string) (types.User, error) {
	res := database.DoQueryRow(db, "SELECT * FROM users WHERE id = ?", id)

	var user types.User
	err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("There is no appointment with this id: %s\n", err)
			return types.User{}, err
		}
		return types.User{}, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user *types.User, id string) (int64, error) {
	res, err := database.DoExec(db, "UPDATE users SET firstName = ?, lastName = ?, email = ?, password = ?, role = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.Password, user.Role, id)
	if err != nil {
		fmt.Printf("Error while updating user: %s\n", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while getting rows affected: %s\n", err)
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteUser(db *sql.DB, id string) (int64, error) {
	//res, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	//res, err := db.Exec("DELETE u FROM users u LEFT JOIN appointments a ON u.id = a.user_id LEFT JOIN shops s on u.id = s.user_id WHERE u.id", id)
	res, err := database.DoExec(db, "DELETE u FROM users u LEFT JOIN shops s on u.id = s.user_id WHERE u.id = ?", id)

	if err != nil {
		fmt.Printf("Error while deleting user: %s\n", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while getting rows affected: %s\n", err)
		return 0, err
	}

	return rowsAffected, nil
}
