package users

import (
	"github.com/itshaydendev/cards/internal"
)

// User represents a user of the service
type User struct {
	ID       int    `json:"id"`
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
		rows.Scan(&user.ID, &user.Username)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetOne finds and returns a user from the database by username
func GetOne(username string) (*User, error) {
	db, err := internal.Database()

	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM users WHERE username = $1 LIMIT 1", username)

	user := User{}

	err = row.Scan(&user.ID, &user.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// Create inserts a new user into the database.
func Create(username string) (*User, error) {
	db, err := internal.Database()

	if err != nil {
		return nil, err
	}

	user := db.QueryRow(`
		INSERT INTO users (
			username
		) VALUES (
			$1
		) RETURNING id, username
	`, username)

	newUser := User{}

	err = user.Scan(&newUser.ID, &newUser.Username)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// Save updates fields in the database where needed
func (u User) Save() error {
	db, err := internal.Database()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		UPDATE users
		SET username = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	stmt.Exec(u.Username, u.ID)

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// TODO implement deleting
