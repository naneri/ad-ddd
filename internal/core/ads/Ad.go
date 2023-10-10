package ads

import "time"

type Ad struct {
	Id        int
	Title     string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
