package repository

import (
	"github.com/lhlyu/iyu/repository/category_repository"
	"github.com/lhlyu/iyu/repository/quanta_repository"
)

type Dao struct {
	*quanta_repository.QuantaDao
	*category_repository.CategoryDao
}

func NewDao(traceId string) *Dao {
	quantaDao := quanta_repository.NewQuantaDao(traceId)
	categoryDao := category_repository.NewCategoryDao(traceId)
	dao := &Dao{
		QuantaDao:   quantaDao,
		CategoryDao: categoryDao,
	}
	return dao
}
