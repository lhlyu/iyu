package repository

import (
	"github.com/lhlyu/iyu/repository/quanta_repository"
)

type Dao struct {
	*quanta_repository.QuantaDao
}

func NewDao(traceId string) *Dao {
	quantaDao := quanta_repository.NewQuantaDao(traceId)
	dao := &Dao{
		QuantaDao: quantaDao,
	}
	return dao
}
