package repository

type Dao struct {
}

func NewDao(traceId string) *Dao {
	dao := &Dao{}
	return dao
}
