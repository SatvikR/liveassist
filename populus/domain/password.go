// Copyright (c) 2021 Satvik Reddy
package domain

import "github.com/matthewhartstonge/argon2"

var argon argon2.Config = argon2.DefaultConfig()

// HashPW hashes a password using the argon2 algorithm
func HashPW(password string) (string, error) {
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

// VerifyPW verifies a hashed password against a raw password
func VerifyPW(password string, encoded string) bool {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(encoded))
	if err != nil {
		return false
	}
	return ok
}
