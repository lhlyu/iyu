package repository

import (
	"github.com/lhlyu/iyu/repository/category_repository"
	"github.com/lhlyu/iyu/repository/quanta_repository"
	"github.com/lhlyu/iyu/repository/tag_repository"
)

type Dao struct {
	QuantaDao   *quanta_repository.Dao
	CategoryDao *category_repository.Dao
	TagDao      *tag_repository.Dao
}

func NewDao(traceId string) *Dao {
	quantaDao := quanta_repository.NewDao(traceId)
	categoryDao := category_repository.NewDao(traceId)
	tagDao := tag_repository.NewDao(traceId)
	dao := &Dao{
		QuantaDao:   quantaDao,
		CategoryDao: categoryDao,
		TagDao:      tagDao,
	}
	return dao
}
