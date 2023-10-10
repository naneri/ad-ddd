package services

import "ad-ddd/internal/core/users"

type CreateUser struct {
	hasher   PasswordHasher
	userRepo users.UserRepo
}

func (cu CreateUser) create(username, email, password string) (users.User, error) {
	user := users.User{
		Username: username,
		Email:    email,
		Password: cu.hasher.Hash(password),
	}

	err := cu.userRepo.Add(user)

	return user, err
}
