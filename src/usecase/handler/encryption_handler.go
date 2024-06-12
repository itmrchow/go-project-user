package handler

type EncryptionHandler interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
