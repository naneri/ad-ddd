package users

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Visits    []Visit `gorm:"-:all"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
