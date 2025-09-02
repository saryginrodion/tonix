package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (*string, error) {
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	generatedHashString := string(generatedHash)
	return &generatedHashString, nil
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
