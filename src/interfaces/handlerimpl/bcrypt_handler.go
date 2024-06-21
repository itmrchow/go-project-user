package handlerimpl

import "golang.org/x/crypto/bcrypt"

type BcryptHandler struct {
}

func NewBcryptHandler() *BcryptHandler {
	return &BcryptHandler{}
}

func (b *BcryptHandler) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (b *BcryptHandler) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (b *BcryptHandler) GetHashingCost(hashedPassword []byte) (int, error) {
	return bcrypt.Cost(hashedPassword)
}
