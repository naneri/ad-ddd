package ads

type AdRepo interface {
	Find(id int) (Ad, error)
	Get(offset, count int) ([]Ad, error)
}
