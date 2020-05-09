package users

import "github.com/itshaydendev/cards/internal"

// User represents a user of the service
type User struct {
	Username string `json:"username"`
}

// GetAll returns a slice of every user in the database
func GetAll(fields string) ([]User, error) {
	users := []User{}

	db, err := internal.Database()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
