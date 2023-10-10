package services

type PasswordHasher interface {
	Hash(password string) string
}
