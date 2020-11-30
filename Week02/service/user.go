package service

import (
	"geek/Week02/dao"
	xerr "github.com/pkg/errors"
)

type User struct {
}

func (u User) GetOneByID(id int) (interface{}, error) {
	data, err := dao.Mysql{}.GetOneByID(id)
	if err != nil {
		return nil, xerr.WithMessagef(err, "req %d", id)
	}

	return data, nil
}
