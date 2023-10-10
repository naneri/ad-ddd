package users

type UserRepo interface {
	Find(Id int) (User, error)
	Add(User) error
	Delete(Id int) error
}
