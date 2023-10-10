package memory

import (
	"ad-ddd/internal/core/users"
	"sync"
	"time"
)

type UserMemoryRepo struct {
	Storage []users.User
	sync.Mutex
}

func (u *UserMemoryRepo) Find(Id int) (users.User, error) {
	if len(u.Storage) <= Id {
		return users.User{}, users.ErrorUserNotFound
	}
	user := u.Storage[Id]

	if user.DeletedAt != nil {
		return users.User{}, users.ErrorUserNotFound
	}

	return u.Storage[Id], nil
}

func (u *UserMemoryRepo) Add(user users.User) error {
	u.Lock()

	user.Id = len(u.Storage)

	u.Storage = append(u.Storage, user)
	u.Unlock()

	return nil
}

func (u *UserMemoryRepo) Delete(Id int) error {
	if len(u.Storage) <= Id {
		return users.ErrorUserNotFound
	}

	user := u.Storage[Id]
	currentTime := time.Now()
	user.DeletedAt = &currentTime

	return nil
}
