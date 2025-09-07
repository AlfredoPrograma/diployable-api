package services

import "golang.org/x/crypto/bcrypt"

type EncryptService interface {
	Encrypt(plain string) (string, error)
	Validate(encrypted string, plain string) bool
}

type encryptService struct {
	encryptCost int
}

func (s encryptService) Encrypt(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), s.encryptCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s encryptService) Validate(encrypted string, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain)); err != nil {
		return false
	}

	return true
}

func NewEncryptService(encryptCost int) EncryptService {
	return encryptService{
		encryptCost: encryptCost,
	}
}
