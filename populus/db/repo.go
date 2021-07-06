// Copyright (c) 2021 Satvik Reddy
package db

import "github.com/go-pg/pg/v10"

// Creates a user and stores it in the database.
func CreateUser(username string, hashedpw string, email string) (int64, string, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, "", err
	}

	newUser := &User{
		Username: username,
		Password: hashedpw,
		Email:    email,
	}
	if _, err := tx.Model(newUser).Insert(); err != nil {
		tx.Rollback()
		return 0, "", err
	}

	if err := tx.Commit(); err != nil {
		return 0, "", err
	}
	return newUser.ID, newUser.Username, nil
}

// FindUserByUsername looks for a user in the database and returns it if found.
func FindUserByUsername(username string) (*User, error) {
	user := new(User)
	if err := db.Model(user).
		Where("? = ?", pg.Ident("username"), username).
		Select(); err != nil {
		return (*User)(nil), err
	}

	return user, nil
}

// FindUserByID finds the user data given a user id
func FindUserByID(userId int64) (*User, error) {
	user := new(User)
	if err := db.Model(user).
		Where("? = ?", pg.Ident("id"), userId).
		Select(); err != nil {
		return (*User)(nil), err
	}

	return user, nil
}
