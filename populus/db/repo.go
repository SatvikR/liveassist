// Copyright (c) 2021 Satvik Reddy
package db

// Creates a user and stores it in the database
func CreateUser(username string, hashedpw string, email string) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	newUser := &user{
		Username: username,
		Password: hashedpw,
		Email:    email,
	}
	if _, err := tx.Model(newUser).Insert(); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return newUser.ID, nil
}
