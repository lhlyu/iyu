package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/errcode"
)

type repositoryError struct {
	fname  string
	sql    string
	params []interface{}
	*errcode.ErrCode
}

func (e *repositoryError) String() string {
	s := ""
	if e.fname != "" {
		s += fmt.Sprintf("%s | ", e.fname)
	}
	if e.sql != "" {
		s += fmt.Sprintf("%s | ", e.sql)
	}
	if len(e.params) > 0 {
		s += fmt.Sprintf("%v | ", e.params)
	}
	if e.Msg != "" {
		s += e.Msg
	}
	return s
}

func NewRepositoryError(fname, sql string, code int, params ...interface{}) *repositoryError {
	errCode := errcode.NewErrcode(code, nil)
	return &repositoryError{
		fname:   "repository." + fname,
		sql:     sql,
		params:  params,
		ErrCode: errCode,
	}
}
