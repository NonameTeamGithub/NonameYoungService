package utilities

import "golang.org/x/crypto/bcrypt"

func CompareHashes(original, hash string) error {
	hashErr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(original))
	return hashErr
}

func MakeHash(value string) (hash string, err error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(value), 8)
	return string(pass), err
}
