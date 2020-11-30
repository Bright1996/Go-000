package dao

import (
	"database/sql"
	"errors"
	xerr "github.com/pkg/errors"
)

var NoFind = errors.New("record not find")

// 简单点直接大写导出 不用初始化依赖
type Mysql struct {
}

func (m Mysql) GetOneByID(id int) (interface{}, error) {
	var err error

	if id == 1 {
		err = sql.ErrNoRows
	}

	if id == 2 {
		err = sql.ErrConnDone
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, xerr.WithStack(NoFind)
	}

	if err != nil {
		return nil, xerr.WithStack(err)
	}

	return id, nil
}
