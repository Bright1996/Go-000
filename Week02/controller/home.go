package controller

import (
	"errors"
	"fmt"
	"geek/Week02/dao"
	"geek/Week02/service"
)

type Home struct {
}

func (h Home) GetHomeInfo(id int) interface{} {
	data, err := service.User{}.GetOneByID(id)
	if err != nil {
		if checkNotFind(err) {
			return fmt.Sprintf("记录不存在")

		}
		return fmt.Sprintf("%+v\n", err)
	}

	return data
}

func checkNotFind(err error) bool {
	return errors.Is(err, dao.NoFind)
}
