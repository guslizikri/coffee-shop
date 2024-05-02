package pkg

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	hasedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hasedBytes), nil
}

func VerifyPass(hashPass, plainPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainPass))
	return err
}
