package core

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword  hash & salt your passwords using the bcrypt package
func HashPassword(plainPwd []byte) (string, error) {
	// Use GenerateFromPassword to hash & salt password.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(plainPwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

// CompareHashAndPassword compares a bcrypt hashed password
// with its possible plaintext equivalent. Returns nil on success,
//  or an error on failure
func CompareHashAndPassword(hashedPwd, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
